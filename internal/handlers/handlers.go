package handlers

import (
	"github.com/ramadoiranedar/user_micro/internal/usecases"
	"github.com/ramadoiranedar/user_micro/internal/utilities"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type handler struct {
	logger   logrus.FieldLogger
	config   *viper.Viper
	usecases usecases.Usecases
}

func NewHandlers(logger logrus.FieldLogger, config *viper.Viper, usecases usecases.Usecases) Handlers {
	utilities.EntryLogTrace(logger, utilities.TraceLog()).Info("setup handlers")

	return &handler{
		logger:   logger,
		config:   config,
		usecases: usecases,
	}
}

type Handlers interface {
	healthHandlers
}
