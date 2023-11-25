package utilities

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type runtimeApp struct {
	config   *viper.Viper
	logger   logrus.FieldLogger
	database *gorm.DB
}

func NewRuntime(configuration *viper.Viper, logger logrus.FieldLogger) (r *runtimeApp) {
	var (
		logEntry = EntryLogTrace(logger, TraceLog())
	)

	logEntry.Info("setup runtime")

	database, err := GetGormDbMysql(configuration, logger)
	if err != nil {
		logEntry.Errorf("database connection error: %v", err)
		os.Exit(1)
	}

	r = &runtimeApp{
		config:   configuration,
		logger:   logger,
		database: database,
	}

	// run migration every start the server
	// r.MigrateDatabase()

	return
}

func (r *runtimeApp) Database() *gorm.DB {
	return r.database
}

func (r *runtimeApp) Logger() logrus.FieldLogger {
	return r.logger
}

func (r *runtimeApp) Config() *viper.Viper {
	return r.config
}

func (r *runtimeApp) Close() {
	EntryLogTrace(r.logger, TraceLog()).
		Infof("%v", "runtime close")
	// TODO: Close at Runtime
}

func (r *runtimeApp) migrateDatabase() {
	EntryLogTrace(r.logger, TraceLog()).
		Info("run migration database")

	r.database.AutoMigrate(
	// TODO: Migration
	// models.User{},
	)
}
