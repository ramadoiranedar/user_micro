package routes

import (
	"github.com/ramadoiranedar/user_micro/gen/restapi/operations"
	"github.com/ramadoiranedar/user_micro/internal/handlers"
	"github.com/ramadoiranedar/user_micro/internal/utilities"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetRoutes(logger logrus.FieldLogger, config *viper.Viper, api *operations.UserMicroServerAPI, apiHandler handlers.Handlers) {
	utilities.EntryLogTrace(logger, utilities.TraceLog()).Info("setup routes")

	// set all route here
	healthRoute(logger, config, api, apiHandler)
}
