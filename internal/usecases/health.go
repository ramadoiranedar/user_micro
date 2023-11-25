package usecases

import (
	"net/http"

	str "github.com/ramadoiranedar/user_micro/internal/constants/strings"
	"github.com/ramadoiranedar/user_micro/internal/utilities"
)

type healthUsecases interface {
	HealthCheckServer() (err error)
}

func (u *usecase) HealthCheckServer() (err error) {
	configKeys := u.config.AllKeys()
	if len(configKeys) < 1 {
		err = utilities.SetError(http.StatusInternalServerError, str.MSG_CONFIG_IS_ERROR)
		return
	}

	if err = u.repositories.HealthCheckDatabase(); err != nil {
		err = utilities.GetError(err)
		return
	}

	return
}
