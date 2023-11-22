package utilities

import (
	"fmt"
	"runtime"

	"github.com/ramadoiranedar/user_micro/internal/constants"
	"github.com/sirupsen/logrus"
)

func InitLogger() logrus.FieldLogger {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = constants.TIME_FORMAT_LOGRUS
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)
	logrus.SetLevel(logrus.DebugLevel)
	CreateLogFields(logrus.StandardLogger(), "utilities-logger", TraceLog()).Info("init logger")

	return logrus.StandardLogger()
}

func CreateLogFields(logger logrus.FieldLogger, moduleName, trace string) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"module": moduleName,
		"trace":  trace,
	})
}

func TraceLog() string {
	pc, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s:%d %s\n", file, line, runtime.FuncForPC(pc).Name())
}
