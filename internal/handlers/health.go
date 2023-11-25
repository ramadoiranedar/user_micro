package handlers

import (
	"net/http"

	"github.com/ramadoiranedar/user_micro/gen/models"
	"github.com/ramadoiranedar/user_micro/gen/restapi/operations/health_check"
	str "github.com/ramadoiranedar/user_micro/internal/constants/strings"
	"github.com/ramadoiranedar/user_micro/internal/utilities"
)

type healthHandlers interface {
	HealthCheckGetV1Health(params health_check.GetV1HealthParams) (response *models.BasicResponse, err error)
}

func (h *handler) HealthCheckGetV1Health(params health_check.GetV1HealthParams) (response *models.BasicResponse, err error) {
	if err = h.usecases.HealthCheckServer(); err != nil {
		err = utilities.GetError(err)
		return
	}

	response = &models.BasicResponse{
		Code:    http.StatusOK,
		Message: str.MSG_OK,
	}

	return
}
