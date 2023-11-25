package repositories

import (
	"github.com/ramadoiranedar/user_micro/internal/utilities"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type repository struct {
	logger   logrus.FieldLogger
	config   *viper.Viper
	database *gorm.DB
}

func NewRepositories(logger logrus.FieldLogger, config *viper.Viper, database *gorm.DB) Repositories {
	utilities.EntryLogTrace(logger, utilities.TraceLog()).Info("setup repositories")

	return &repository{
		logger:   logger,
		config:   config,
		database: database,
	}
}

type Repositories interface {
	healthRepositories
}
