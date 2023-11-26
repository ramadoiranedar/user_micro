package utilities

import (
	"github.com/ramadoiranedar/user_micro/gen/restapi/operations"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AuthorizationDTO struct {
	Logger logrus.FieldLogger
	Config *viper.Viper
	Api    *operations.UserMicroServerAPI
}

func NewAuthorization(authorizationDTO AuthorizationDTO) *AuthorizationDTO {
	EntryLogTrace(authorizationDTO.Logger, TraceLog()).Info("setup authorization")

	return &authorizationDTO
}
