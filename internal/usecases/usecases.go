package usecases

import (
	"github.com/ramadoiranedar/user_micro/internal/repositories"
	"github.com/ramadoiranedar/user_micro/internal/utilities"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type UsecasesDTO struct {
	Logger       logrus.FieldLogger
	Config       *viper.Viper
	Repositories repositories.Repositories
}

func NewUsecases(usecasesDTO UsecasesDTO) Usecases {
	utilities.EntryLogTrace(usecasesDTO.Logger, utilities.TraceLog()).Info("setup usecases")

	return &usecasesDTO
}

type Usecases interface {
	healthUsecases
}
