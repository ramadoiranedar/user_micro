package handlers

import (
	"net/http"

	"github.com/ramadoiranedar/user_micro/gen/models"
	"github.com/ramadoiranedar/user_micro/gen/restapi/operations/health_check"
	str "github.com/ramadoiranedar/user_micro/internal/constants/strings"
	"github.com/ramadoiranedar/user_micro/internal/utilities"
)

type healthHandlers interface {
	HealthCheckGetV1Health(params health_check.GetV1HealthParams) (response *models.HealthCheckResponse, err error)
}

func (h *handler) HealthCheckGetV1Health(params health_check.GetV1HealthParams) (response *models.HealthCheckResponse, err error) {
	results, err := h.usecases.HealthCheckServer()
	if err != nil {
		err = utilities.GetError(err)
		return
	}

	response = &models.HealthCheckResponse{
		Code:    http.StatusOK,
		Message: str.MSG_OK,
		Results: results,
	}

	return
}
