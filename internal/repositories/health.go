package repositories

import (
	"net/http"

	"github.com/ramadoiranedar/user_micro/internal/utilities"
)

type healthRepositories interface {
	HealthCheckDatabase() (err error)
}

func (r *RepositoriesDTO) HealthCheckDatabase() (err error) {
	var dbVersion *string
	if err = r.Database.Raw("SELECT VERSION() AS version;").Scan(&dbVersion).Error; err != nil {
		err = utilities.SetError(http.StatusNotFound, err.Error())
		return
	}

	return
}
