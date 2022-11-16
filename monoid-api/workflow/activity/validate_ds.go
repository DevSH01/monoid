package activity

import (
	"context"
	"encoding/json"

	"github.com/brist-ai/monoid/model"
	"github.com/brist-ai/monoid/monoidprotocol"
	"github.com/brist-ai/monoid/monoidprotocol/docker"
	"go.temporal.io/sdk/activity"
)

type ValidateDSArgs struct {
	SiloSpecID string
	Config     []byte
}

func (a *Activity) ValidateDataSiloDef(ctx context.Context, args ValidateDSArgs) (*monoidprotocol.MonoidValidateMessage, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("Validating silo def")

	spec := model.SiloSpecification{}
	if err := a.Conf.DB.Where("id = ?", args.SiloSpecID).First(&spec).Error; err != nil {
		logger.Error("Could not find silo spec: %v", err)
		return nil, err
	}

	mp, err := docker.NewDockerMP(spec.DockerImage, spec.DockerTag)
	if err != nil {
		logger.Error("Error creating docker client: %v", err)
		return nil, err
	}

	defer mp.Teardown(ctx)

	if err := mp.InitConn(ctx); err != nil {
		logger.Error("Error creating docker connection: %v", err)
		return nil, err
	}

	confString := model.SecretString("")
	confString.Scan(args.Config)

	conf := map[string]interface{}{}
	json.Unmarshal([]byte(confString), &conf)

	logger.Info("validating")

	validate, err := mp.Validate(ctx, conf)

	if err != nil {
		logger.Error("Error running validate: %v", err)
		return nil, err
	}

	return validate, nil
}
