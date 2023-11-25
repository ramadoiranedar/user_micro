package repositories

import (
	"net/http"

	"github.com/ramadoiranedar/user_micro/internal/utilities"
)

type healthRepositories interface {
	HealthCheckDatabase() (err error)
}

func (r *repository) HealthCheckDatabase() (err error) {
	var dbVersion *string
	if err = r.database.Raw("SELECT VERSION() AS version;").Scan(&dbVersion).Error; err != nil {
		err = utilities.SetError(http.StatusNotFound, err.Error())
		return
	}

	return
}
