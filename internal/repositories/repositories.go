package repositories

import (
	"github.com/ramadoiranedar/user_micro/internal/utilities"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type RepositoriesDTO struct {
	Logger   logrus.FieldLogger
	Config   *viper.Viper
	Database *gorm.DB
}

func NewRepositories(repositoriesDTO RepositoriesDTO) Repositories {
	utilities.EntryLogTrace(repositoriesDTO.Logger, utilities.TraceLog()).Info("setup repositories")

	return &repositoriesDTO
}

type Repositories interface {
	healthRepositories
}
