package routes

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ramadoiranedar/user_micro/gen/models"
	"github.com/ramadoiranedar/user_micro/gen/restapi/operations/health_check"
	"github.com/ramadoiranedar/user_micro/internal/utilities"
)

func healthRoute(routesDTO RoutesDTO) {

	/*
	* Baseline code structure route
	*	Please, following this code styles correctly
	* make sure you when create function using DTO object from the handler layer
	* and so on until the last layer repository. (eg. h.HealthCheckGetV1Health
	* (params) using DTO health_check.GetV1HealthParams)
	* and you can passing everything you need from route params above to handler
	**/

	// GET /api/v1/health
	routesDTO.Api.HealthCheckGetV1HealthHandler = health_check.GetV1HealthHandlerFunc(func(params health_check.GetV1HealthParams) middleware.Responder {
		logEntry := utilities.EntryLogTraceEndpoint(routesDTO.Logger, routesDTO.Config, params.HTTPRequest, utilities.TraceLog())

		response, err := routesDTO.Handlers.HealthCheckGetV1Health(params)
		if err != nil {
			errResponse := utilities.GetError(err)
			logEntry.Error(errResponse.Code(), errResponse)
			return health_check.NewGetV1HealthDefault(int(errResponse.Code())).WithPayload(&models.BasicResponse{
				Code:    errResponse.Code(),
				Message: errResponse.Error(),
			})
		}

		logEntry.Info(response)
		return health_check.NewGetV1HealthOK().WithPayload(response)
	})

	// more route here
	// ...

}
