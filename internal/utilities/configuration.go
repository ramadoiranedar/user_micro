package utilities

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var FLAGS = struct {
	AppConfig string `long:"config" description:"main application configuration YAML path"`
}{}

func InitConfiguration(logger logrus.FieldLogger) *viper.Viper {
	var (
		logFields = CreateLogFields(logrus.StandardLogger(), "utilities-configuration", TraceLog())
	)
	logFields.Info("init configuration")

	cfg := viper.New()
	cfg.SetConfigName(filepath.Base(FLAGS.AppConfig))
	cfg.SetConfigType("yaml")
	cfg.AddConfigPath(filepath.Dir(FLAGS.AppConfig))
	cfg.AddConfigPath("./config/")
	cfg.AddConfigPath("./etc/")
	cfg.AddConfigPath("./")

	err := cfg.ReadInConfig()
	if err != nil {
		logFields.Errorf("invalid app config at %s", FLAGS.AppConfig)
		os.Exit(1)
	}

	logFields.Info("configuration app is ok")
	return cfg
}
