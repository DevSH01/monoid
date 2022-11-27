package requestactivity

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/brist-ai/monoid/model"
	"github.com/brist-ai/monoid/monoidprotocol"
	monoidactivity "github.com/brist-ai/monoid/workflow/activity"
	"github.com/google/uuid"
	"go.temporal.io/sdk/activity"
)

type ProcessRequestArgs struct {
	ProtocolRequestStatus []monoidprotocol.MonoidRequestStatus
	RequestStatusIDs      []string
}

type ProcessRequestItem struct {
	Error           *RequestStatusError
	RequestStatusID string
}

type ProcessRequestResult struct {
	ResultItems []ProcessRequestItem
}

func (a *RequestActivity) ProcessRequestResults(
	ctx context.Context,
	args ProcessRequestArgs,
) (ProcessRequestResult, error) {
	logger := activity.GetLogger(ctx)
	resultMap := map[string]ProcessRequestItem{}
	requestStatuses := []*model.RequestStatus{}
	protocolStatusMap := map[monoidactivity.DataSourceMatcher]monoidprotocol.MonoidRequestStatus{}
	matcherRequestStatusMap := map[monoidactivity.DataSourceMatcher]model.RequestStatus{}

	if err := a.Conf.DB.Model(model.RequestStatus{}).
		Preload("DataSource").
		Preload("DataSource.SiloDefinition").
		Preload("DataSource.SiloDefinition.SiloSpecification").
		Preload("Request").
		Where("id IN ?", args.RequestStatusIDs).First(&requestStatuses).Error; err != nil {
		return ProcessRequestResult{}, err
	}

	for _, prs := range args.ProtocolRequestStatus {
		protocolStatusMap[monoidactivity.NewDataSourceMatcher(
			prs.SchemaName,
			prs.SchemaGroup,
		)] = prs
	}

	var siloDef *model.SiloDefinition
	var request *model.Request

	handles := make([]monoidprotocol.MonoidRequestHandle, 0, len(requestStatuses))

	for _, rs := range requestStatuses {
		dsm := monoidactivity.NewDataSourceMatcher(
			rs.DataSource.Name,
			rs.DataSource.Group,
		)
		prstat, ok := protocolStatusMap[dsm]

		requestSilo := rs.DataSource.SiloDefinition
		if siloDef != nil && siloDef.ID != requestSilo.ID {
			return ProcessRequestResult{}, fmt.Errorf("all requests must be for the same silo")
		} else if siloDef == nil {
			siloDef = &requestSilo
		}

		rsRequest := rs.Request
		if request != nil && request.ID != rsRequest.ID {
			return ProcessRequestResult{}, fmt.Errorf("all requests must be for the same request")
		} else if request == nil {
			request = &rsRequest
		}

		if !ok {
			resultMap[rs.ID] = ProcessRequestItem{
				Error: &RequestStatusError{Message: "request status not provided"},
			}

			continue
		}

		if prstat.RequestStatus != monoidprotocol.MonoidRequestStatusRequestStatusCOMPLETE {
			resultMap[rs.ID] = ProcessRequestItem{
				Error: &RequestStatusError{Message: "request can only be read when the status is complete"},
			}

			continue
		}

		dataType := prstat.DataType
		if dataType == nil || *dataType == monoidprotocol.MonoidRequestStatusDataTypeNONE {
			resultMap[rs.ID] = ProcessRequestItem{}
			continue
		}

		handle := monoidprotocol.MonoidRequestHandle{}
		if err := json.Unmarshal([]byte(rs.RequestHandle), &handle); err != nil {
			resultMap[rs.ID] = ProcessRequestItem{
				Error: &RequestStatusError{Message: err.Error()},
			}
		}

		handles = append(handles, handle)
		matcherRequestStatusMap[dsm] = *rs
	}

	if siloDef == nil {
		return ProcessRequestResult{}, nil
	}

	if len(handles) > 0 {
		siloSpec := siloDef.SiloSpecification

		conf := map[string]interface{}{}
		if err := json.Unmarshal([]byte(siloDef.Config), &conf); err != nil {
			return ProcessRequestResult{}, err
		}

		// Create a temporary directory that can be used by the docker container
		dir, err := ioutil.TempDir("/tmp/monoid", "monoid")
		if err != nil {
			return ProcessRequestResult{}, err
		}

		defer os.RemoveAll(dir)

		// Start the docker protocol
		protocol, err := a.Conf.ProtocolFactory.NewMonoidProtocol(
			siloSpec.DockerImage, siloSpec.DockerTag, dir,
		)
		if err != nil {
			return ProcessRequestResult{}, err
		}

		defer protocol.Teardown(ctx)

		logChan, err := protocol.AttachLogs(ctx)
		if err != nil {
			return ProcessRequestResult{}, err
		}

		go func() {
			for l := range logChan {
				logger.Debug(l.Message)
			}
		}()

		recordCh, err := protocol.RequestResults(ctx, conf, monoidprotocol.MonoidRequestsMessage{
			Handles: handles,
		})

		if err != nil {
			return ProcessRequestResult{}, err
		}

		type queryResult struct {
			resultType model.ResultType
			data       any
		}

		queryResults := map[string]*queryResult{}

		for record := range recordCh {
			dsm := monoidactivity.NewDataSourceMatcher(
				record.SchemaName,
				record.SchemaGroup,
			)

			rs, ok := matcherRequestStatusMap[dsm]

			if !ok {
				logger.Warn("Unknown data source found", record.SchemaName, record.SchemaGroup)
				continue
			}

			prs, ok := protocolStatusMap[dsm]
			if !ok {
				logger.Warn("Unknown data source found", record.SchemaName, record.SchemaGroup)
				continue
			}

			dataType := prs.DataType
			switch *dataType {
			case monoidprotocol.MonoidRequestStatusDataTypeFILE:
				_, ok := queryResults[rs.ID]
				if !ok {
					queryResults[rs.ID] = &queryResult{
						resultType: model.ResultTypeFile,
						data:       map[string]interface{}{},
					}
				}

				data, ok := queryResults[rs.ID].data.(map[string]interface{})
				if !ok {
					logger.Warn("Error casting existing data")
					continue
				}

				if len(data) != 0 {
					logger.Warn("File data results should only be one file path, got multiple.")
				}

				if record.File == nil {
					logger.Warn("File attr must not be nil")
					continue
				}

				queryResults[rs.ID].data = map[string]interface{}{
					"filePath": *record.File,
				}
			case monoidprotocol.MonoidRequestStatusDataTypeRECORDS:
				_, ok := queryResults[rs.ID]
				if !ok {
					queryResults[rs.ID] = &queryResult{
						resultType: model.ResultTypeRecordsJSON,
						data:       []*monoidprotocol.MonoidRecordData{},
					}
				}

				data, ok := queryResults[rs.ID].data.([]*monoidprotocol.MonoidRecordData)
				if !ok {
					logger.Warn("Error casting existing data")
					continue
				}

				queryResults[rs.ID].data = append(data, &record.Data)
			}
		}

		// Write the records back to the db
		for rsID, qr := range queryResults {
			if request.Type == model.UserDataRequestTypeQuery {
				records, err := json.Marshal(qr.data)
				if err != nil {
					resultMap[rsID] = ProcessRequestItem{Error: &RequestStatusError{
						Message: err.Error(),
					}}

					continue
				}

				r := model.SecretString(records)

				if err = a.Conf.DB.Create(&model.QueryResult{
					ID:              uuid.NewString(),
					RequestStatusID: rsID,
					Records:         &r,
					ResultType:      qr.resultType,
				}).Error; err != nil {
					resultMap[rsID] = ProcessRequestItem{Error: &RequestStatusError{
						Message: err.Error(),
					}}

					continue
				}
			}

			resultMap[rsID] = ProcessRequestItem{}
		}
	}

	results := make([]ProcessRequestItem, len(args.RequestStatusIDs))
	for i, s := range args.RequestStatusIDs {
		res, ok := resultMap[s]
		if !ok {
			res = ProcessRequestItem{
				Error: &RequestStatusError{Message: fmt.Sprintf("could not find status for %s", s)},
			}
		}

		res.RequestStatusID = s
		results[i] = res
	}

	return ProcessRequestResult{ResultItems: results}, nil
}
