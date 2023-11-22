package utilities

import (
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

/**
*	@Cfg as viper config
* @Log as logrus logger
* @Db as gorm DB
 */
type Runtime struct {
	Cfg *viper.Viper
	Log logrus.FieldLogger
	Db  *gorm.DB
}

func InitRuntime() (runtime *Runtime) {
	var (
		logger    = InitLogger()
		config    = InitConfiguration(logger)
		logFields = CreateLogFields(logger, "utilities-runtime", TraceLog())
	)

	logFields.Info("run runtime")

	database, err := SetGormDbMysql(config, logger)
	if err != nil {
		logFields.Errorf("database connection error: %v", err)
		os.Exit(1)
	}

	runtime = &Runtime{
		Cfg: config,
		Log: logger,
		Db:  database,
	}

	return
}

func (rt *Runtime) Database() *gorm.DB {
	return rt.Db
}

func (rt *Runtime) Logger() logrus.FieldLogger {
	return rt.Log
}

func (rt *Runtime) Config() *viper.Viper {
	return rt.Cfg
}

func (rt *Runtime) RunMigration() {
	CreateLogFields(rt.Log, "utilities-runtime", TraceLog()).
		Info("run migration database")

	rt.Db.AutoMigrate(
	// TODO: Migration
	// models.User{},
	)
}

func (rt *Runtime) SetError(code int, msg string, args ...interface{}) error {
	CreateLogFields(rt.Log, "utilities-runtime", TraceLog()).
		Infof("%v", msg)

	return errors.New(int32(code), msg)
}

func (rt *Runtime) GetError(err error) errors.Error {
	CreateLogFields(rt.Log, "utilities-runtime", TraceLog()).
		Errorf("%v", err)
	if v, ok := err.(errors.Error); ok {
		return v
	}

	return errors.New(http.StatusInternalServerError, err.Error())
}

func (rt *Runtime) Close() {
	CreateLogFields(rt.Log, "utilities-runtime", TraceLog()).
		Infof("%v", "runtime close")
	// TODO: Close
}
