package user_micro

import (
	"os"

	"github.com/ramadoiranedar/user_micro/internal/utilities"
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
		logEntry = utilities.EntryLogTrace(logger, utilities.TraceLog())
	)

	logEntry.Info("setup runtime")

	database, err := utilities.GetGormDbMysql(configuration, logger)
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
	utilities.EntryLogTrace(r.logger, utilities.TraceLog()).
		Infof("%v", "runtime close")
	// TODO: Close at Runtime
}

func (r *runtimeApp) migrateDatabase() {
	utilities.EntryLogTrace(r.logger, utilities.TraceLog()).
		Info("run migration database")

	r.database.AutoMigrate(
	// TODO: Migration
	// models.User{},
	)
}
