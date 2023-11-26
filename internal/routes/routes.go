package routes

import (
	"github.com/ramadoiranedar/user_micro/gen/restapi/operations"
	"github.com/ramadoiranedar/user_micro/internal/handlers"
	"github.com/ramadoiranedar/user_micro/internal/utilities"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type RoutesDTO struct {
	Logger   logrus.FieldLogger
	Config   *viper.Viper
	Api      *operations.UserMicroServerAPI
	Handlers handlers.Handlers
}

func SetRoutes(routesDTO RoutesDTO) {
	utilities.EntryLogTrace(routesDTO.Logger, utilities.TraceLog()).Info("setup routes")

	// set all route here
	healthRoute(routesDTO)
}
