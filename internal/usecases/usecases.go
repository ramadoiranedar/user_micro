package usecases

import (
	"github.com/ramadoiranedar/user_micro/internal/repositories"
	"github.com/ramadoiranedar/user_micro/internal/utilities"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type usecase struct {
	logger       logrus.FieldLogger
	config       *viper.Viper
	repositories repositories.Repositories
}

func NewUsecases(logger logrus.FieldLogger, config *viper.Viper, repositories repositories.Repositories) Usecases {
	utilities.EntryLogTrace(logger, utilities.TraceLog()).Info("setup usecases")

	return &usecase{
		logger:       logger,
		config:       config,
		repositories: repositories,
	}
}

type Usecases interface {
	healthUsecases
}
