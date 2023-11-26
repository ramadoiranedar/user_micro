package usecases

import (
	"net/http"

	"github.com/ramadoiranedar/user_micro/gen/models"
	"github.com/ramadoiranedar/user_micro/internal/constants"
	str "github.com/ramadoiranedar/user_micro/internal/constants/strings"
	"github.com/ramadoiranedar/user_micro/internal/utilities"
)

type healthUsecases interface {
	HealthCheckServer() (results *models.HealthCheckResponseResults, err error)
}

func (u *UsecasesDTO) HealthCheckServer() (results *models.HealthCheckResponseResults, err error) {
	configKeys := u.Config.AllKeys()
	if len(configKeys) < 1 {
		err = utilities.SetError(http.StatusInternalServerError, str.MSG_CONFIG_IS_ERROR)
		return
	}

	if err = u.Repositories.HealthCheckDatabase(); err != nil {
		err = utilities.GetError(err)
		return
	}

	constants := constants.NewConstants(u.Config)

	results = &models.HealthCheckResponseResults{
		AppVersion:     constants.GetAppVersion(),
		AppName:        constants.GetAppName(),
		AppDescription: constants.GetAppDescription(),
		AppCopyright:   constants.GetAppCopyright(),
		AppEnvironment: constants.GetAppEnvironment(),
		AppMaxUploadMb: constants.GetAppMaxUploadMB(),
	}

	return
}
