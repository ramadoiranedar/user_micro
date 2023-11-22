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
	CreateLogFields(logrus.StandardLogger(), TraceLog()).Info("setup logger")

	return logrus.StandardLogger()
}

func CreateLogFields(logger logrus.FieldLogger, trace string) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"trace": trace,
	})
}

func TraceLog() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}
