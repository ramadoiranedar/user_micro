package utilities

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewConfiguration(logger logrus.FieldLogger, flagAppConfig string) *viper.Viper {
	var (
		logEntry = EntryLogTrace(logger, TraceLog())
	)
	logEntry.Infof("setup configuration from %s", flagAppConfig)

	config := viper.New()
	config.SetConfigName(filepath.Base(flagAppConfig))
	config.SetConfigType("yaml")
	config.AddConfigPath(filepath.Dir(flagAppConfig))
	config.AddConfigPath("./")
	config.AddConfigPath("./etc/")
	config.AddConfigPath("./configs/")

	err := config.ReadInConfig()
	if err != nil {
		logEntry.Error(err)
		os.Exit(1)
	}

	return config
}
