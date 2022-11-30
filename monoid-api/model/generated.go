// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type CategoryQuery struct {
	AnyCategory *bool    `json:"anyCategory"`
	NoCategory  *bool    `json:"noCategory"`
	CategoryIDs []string `json:"categoryIDs"`
}

type CreateCategoryInput struct {
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceID"`
}

type CreateDataSourceInput struct {
	SiloDefinitionID string   `json:"siloDefinitionID"`
	Description      *string  `json:"description"`
	PropertyIDs      []string `json:"propertyIDs"`
}

type CreatePropertyInput struct {
	CategoryIDs  []string `json:"categoryIDs"`
	DataSourceID string   `json:"dataSourceID"`
	PurposeIDs   []string `json:"purposeIDs"`
}

type CreateSiloDefinitionInput struct {
	Description         *string  `json:"description"`
	SiloSpecificationID string   `json:"siloSpecificationID"`
	WorkspaceID         string   `json:"workspaceID"`
	SubjectIDs          []string `json:"subjectIDs"`
	SiloData            *string  `json:"siloData"`
	Name                string   `json:"name"`
}

type CreateSiloSpecificationInput struct {
	Name        string  `json:"name"`
	WorkspaceID string  `json:"workspaceID"`
	LogoURL     *string `json:"logoURL"`
	DockerImage string  `json:"dockerImage"`
	Schema      *string `json:"schema"`
}

type CreateSubjectInput struct {
	Name        string `json:"name"`
	WorkspaceID string `json:"workspaceID"`
}

type CreateUserPrimaryKeyInput struct {
	Name          string `json:"name"`
	APIIdentifier string `json:"apiIdentifier"`
	WorkspaceID   string `json:"workspaceId"`
}

type CreateWorkspaceInput struct {
	Name     string    `json:"name"`
	Settings []*KVPair `json:"settings"`
}

type DataDiscoveriesListResult struct {
	Discoveries    []*DataDiscovery `json:"discoveries"`
	NumDiscoveries int              `json:"numDiscoveries"`
}

type DataMapQuery struct {
	Categories      *CategoryQuery `json:"categories"`
	SiloDefinitions []string       `json:"siloDefinitions"`
}

type DataMapResult struct {
	DataMapRows []*DataMapRow `json:"dataMapRows"`
	NumRows     int           `json:"numRows"`
}

type DownloadLink struct {
	URL string `json:"url"`
}

type HandleAllDiscoveriesInput struct {
	SiloID string          `json:"siloId"`
	Action DiscoveryAction `json:"action"`
}

type HandleDiscoveryInput struct {
	DiscoveryID string          `json:"discoveryId"`
	Action      DiscoveryAction `json:"action"`
}

type JobsResult struct {
	Jobs    []*Job `json:"jobs"`
	NumJobs int    `json:"numJobs"`
}

type KVPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MonoidRecordResponse struct {
	Data        string  `json:"data"`
	SchemaGroup *string `json:"SchemaGroup"`
	SchemaName  string  `json:"SchemaName"`
}

type RequestStatusListResult struct {
	RequestStatusRows []*RequestStatus `json:"requestStatusRows"`
	NumStatuses       int              `json:"numStatuses"`
}

type RequestStatusQuery struct {
	SiloDefinitions []string `json:"siloDefinitions"`
}

type RequestsResult struct {
	Requests    []*Request `json:"requests"`
	NumRequests int        `json:"numRequests"`
}

type UpdateCategoryInput struct {
	Name *string `json:"name"`
}

type UpdateDataSourceInput struct {
	ID          string  `json:"id"`
	Description *string `json:"description"`
}

type UpdatePropertyInput struct {
	ID          string   `json:"id"`
	CategoryIDs []string `json:"categoryIDs"`
	PurposeIDs  []string `json:"purposeIDs"`
}

type UpdateSiloDefinitionInput struct {
	ID          string   `json:"id"`
	WorkspaceID string   `json:"workspaceId"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	SubjectIDs  []string `json:"subjectIDs"`
	SiloData    *string  `json:"siloData"`
}

type UpdateSiloSpecificationInput struct {
	ID          string  `json:"id"`
	DockerImage *string `json:"dockerImage"`
	Schema      *string `json:"schema"`
	Name        *string `json:"name"`
	LogoURL     *string `json:"logoUrl"`
}

type UpdateSubjectInput struct {
	Name *string `json:"name"`
}

type UpdateUserPrimaryKeyInput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateWorkspaceSettingsInput struct {
	WorkspaceID string    `json:"workspaceID"`
	Settings    []*KVPair `json:"settings"`
}

type UserDataRequestInput struct {
	PrimaryKeys []*UserPrimaryKeyInput `json:"primaryKeys"`
	WorkspaceID string                 `json:"workspaceId"`
	Type        UserDataRequestType    `json:"type"`
}

type UserPrimaryKeyInput struct {
	APIIdentifier string `json:"apiIdentifier"`
	Value         string `json:"value"`
}

type DiscoveryAction string

const (
	DiscoveryActionAccept DiscoveryAction = "ACCEPT"
	DiscoveryActionReject DiscoveryAction = "REJECT"
)

var AllDiscoveryAction = []DiscoveryAction{
	DiscoveryActionAccept,
	DiscoveryActionReject,
}

func (e DiscoveryAction) IsValid() bool {
	switch e {
	case DiscoveryActionAccept, DiscoveryActionReject:
		return true
	}
	return false
}

func (e DiscoveryAction) String() string {
	return string(e)
}

func (e *DiscoveryAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DiscoveryAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DiscoveryAction", str)
	}
	return nil
}

func (e DiscoveryAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DiscoveryStatus string

const (
	DiscoveryStatusOpen     DiscoveryStatus = "OPEN"
	DiscoveryStatusAccepted DiscoveryStatus = "ACCEPTED"
	DiscoveryStatusRejected DiscoveryStatus = "REJECTED"
)

var AllDiscoveryStatus = []DiscoveryStatus{
	DiscoveryStatusOpen,
	DiscoveryStatusAccepted,
	DiscoveryStatusRejected,
}

func (e DiscoveryStatus) IsValid() bool {
	switch e {
	case DiscoveryStatusOpen, DiscoveryStatusAccepted, DiscoveryStatusRejected:
		return true
	}
	return false
}

func (e DiscoveryStatus) String() string {
	return string(e)
}

func (e *DiscoveryStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DiscoveryStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DiscoveryStatus", str)
	}
	return nil
}

func (e DiscoveryStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DiscoveryType string

const (
	DiscoveryTypeDataSourceMissing DiscoveryType = "DATA_SOURCE_MISSING"
	DiscoveryTypeDataSourceFound   DiscoveryType = "DATA_SOURCE_FOUND"
	DiscoveryTypePropertyFound     DiscoveryType = "PROPERTY_FOUND"
	DiscoveryTypePropertyMissing   DiscoveryType = "PROPERTY_MISSING"
	DiscoveryTypeCategoryFound     DiscoveryType = "CATEGORY_FOUND"
)

var AllDiscoveryType = []DiscoveryType{
	DiscoveryTypeDataSourceMissing,
	DiscoveryTypeDataSourceFound,
	DiscoveryTypePropertyFound,
	DiscoveryTypePropertyMissing,
	DiscoveryTypeCategoryFound,
}

func (e DiscoveryType) IsValid() bool {
	switch e {
	case DiscoveryTypeDataSourceMissing, DiscoveryTypeDataSourceFound, DiscoveryTypePropertyFound, DiscoveryTypePropertyMissing, DiscoveryTypeCategoryFound:
		return true
	}
	return false
}

func (e DiscoveryType) String() string {
	return string(e)
}

func (e *DiscoveryType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DiscoveryType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DiscoveryType", str)
	}
	return nil
}

func (e DiscoveryType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FullRequestStatus string

const (
	FullRequestStatusCreated       FullRequestStatus = "CREATED"
	FullRequestStatusInProgress    FullRequestStatus = "IN_PROGRESS"
	FullRequestStatusExecuted      FullRequestStatus = "EXECUTED"
	FullRequestStatusPartialFailed FullRequestStatus = "PARTIAL_FAILED"
	FullRequestStatusFailed        FullRequestStatus = "FAILED"
)

var AllFullRequestStatus = []FullRequestStatus{
	FullRequestStatusCreated,
	FullRequestStatusInProgress,
	FullRequestStatusExecuted,
	FullRequestStatusPartialFailed,
	FullRequestStatusFailed,
}

func (e FullRequestStatus) IsValid() bool {
	switch e {
	case FullRequestStatusCreated, FullRequestStatusInProgress, FullRequestStatusExecuted, FullRequestStatusPartialFailed, FullRequestStatusFailed:
		return true
	}
	return false
}

func (e FullRequestStatus) String() string {
	return string(e)
}

func (e *FullRequestStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FullRequestStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FullRequestStatus", str)
	}
	return nil
}

func (e FullRequestStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type JobStatus string

const (
	JobStatusQueued        JobStatus = "QUEUED"
	JobStatusRunning       JobStatus = "RUNNING"
	JobStatusCompleted     JobStatus = "COMPLETED"
	JobStatusPartialFailed JobStatus = "PARTIAL_FAILED"
	JobStatusFailed        JobStatus = "FAILED"
)

var AllJobStatus = []JobStatus{
	JobStatusQueued,
	JobStatusRunning,
	JobStatusCompleted,
	JobStatusPartialFailed,
	JobStatusFailed,
}

func (e JobStatus) IsValid() bool {
	switch e {
	case JobStatusQueued, JobStatusRunning, JobStatusCompleted, JobStatusPartialFailed, JobStatusFailed:
		return true
	}
	return false
}

func (e JobStatus) String() string {
	return string(e)
}

func (e *JobStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = JobStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid JobStatus", str)
	}
	return nil
}

func (e JobStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type RequestStatusType string

const (
	RequestStatusTypeCreated    RequestStatusType = "CREATED"
	RequestStatusTypeInProgress RequestStatusType = "IN_PROGRESS"
	RequestStatusTypeExecuted   RequestStatusType = "EXECUTED"
	RequestStatusTypeFailed     RequestStatusType = "FAILED"
)

var AllRequestStatusType = []RequestStatusType{
	RequestStatusTypeCreated,
	RequestStatusTypeInProgress,
	RequestStatusTypeExecuted,
	RequestStatusTypeFailed,
}

func (e RequestStatusType) IsValid() bool {
	switch e {
	case RequestStatusTypeCreated, RequestStatusTypeInProgress, RequestStatusTypeExecuted, RequestStatusTypeFailed:
		return true
	}
	return false
}

func (e RequestStatusType) String() string {
	return string(e)
}

func (e *RequestStatusType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = RequestStatusType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RequestStatusType", str)
	}
	return nil
}

func (e RequestStatusType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserDataRequestType string

const (
	UserDataRequestTypeDelete UserDataRequestType = "DELETE"
	UserDataRequestTypeQuery  UserDataRequestType = "QUERY"
)

var AllUserDataRequestType = []UserDataRequestType{
	UserDataRequestTypeDelete,
	UserDataRequestTypeQuery,
}

func (e UserDataRequestType) IsValid() bool {
	switch e {
	case UserDataRequestTypeDelete, UserDataRequestTypeQuery:
		return true
	}
	return false
}

func (e UserDataRequestType) String() string {
	return string(e)
}

func (e *UserDataRequestType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserDataRequestType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserDataRequestType", str)
	}
	return nil
}

func (e UserDataRequestType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
