package utilities

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfiguration(logger logrus.FieldLogger, flagAppConfig string) *viper.Viper {
	var (
		logFields = CreateLogFields(logrus.StandardLogger(), TraceLog())
	)
	logFields.Infof("setup configuration from %s", flagAppConfig)

	config := viper.New()
	config.SetConfigName(filepath.Base(flagAppConfig))
	config.SetConfigType("yaml")
	config.AddConfigPath(filepath.Dir(flagAppConfig))
	config.AddConfigPath("./")
	config.AddConfigPath("./etc/")
	config.AddConfigPath("./configs/")

	err := config.ReadInConfig()
	if err != nil {
		logFields.Error(err)
		os.Exit(1)
	}

	return config
}
