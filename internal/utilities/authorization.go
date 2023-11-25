package utilities

import (
	"github.com/ramadoiranedar/user_micro/gen/restapi/operations"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type authorization struct {
	logger logrus.FieldLogger
	config *viper.Viper
	api    *operations.UserMicroServerAPI
}

func NewAuthorization(logger logrus.FieldLogger, config *viper.Viper, api *operations.UserMicroServerAPI) *authorization {
	EntryLogTrace(logger, TraceLog()).Info("setup authorization")

	return &authorization{
		logger: logger,
		config: config,
		api:    api,
	}
}
