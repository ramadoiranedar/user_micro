package handlers

import (
	"github.com/ramadoiranedar/user_micro/internal/usecases"
	"github.com/ramadoiranedar/user_micro/internal/utilities"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type HandlersDTO struct {
	Logger   logrus.FieldLogger
	Config   *viper.Viper
	Usecases usecases.Usecases
}

func NewHandlers(handlersDTO HandlersDTO) Handlers {
	utilities.EntryLogTrace(handlersDTO.Logger, utilities.TraceLog()).Info("setup handlers")

	return &handlersDTO
}

type Handlers interface {
	healthHandlers
}
