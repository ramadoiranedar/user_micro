package utilities

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/ramadoiranedar/user_micro/internal/constants"
	str "github.com/ramadoiranedar/user_micro/internal/constants/strings"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewLogger() logrus.FieldLogger {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = str.TIME_FORMAT_LOGRUS
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)
	logrus.SetLevel(logrus.DebugLevel)
	EntryLogTrace(logrus.StandardLogger(), TraceLog()).Info("setup logger")

	return logrus.StandardLogger()
}

func TraceLog() string {
	pc, _, _, _ := runtime.Caller(1)
	return fmt.Sprintf("%s", runtime.FuncForPC(pc).Name())
}

func EntryLogTrace(logger logrus.FieldLogger, trace interface{}) *logrus.Entry {
	return logger.WithField(str.FIELD_TRACE, trace)
}

func EntryLogTraceEndpoint(logger logrus.FieldLogger, config *viper.Viper, request *http.Request, trace interface{}) *logrus.Entry {
	constants := constants.NewConstants(config)

	return EntryLogTrace(logrus.StandardLogger(), trace).WithFields(logrus.Fields{
		str.FIELD_TRACE:        trace,
		str.FIELD_ENDPOINT:     fmt.Sprintf("%s %s", request.Method, request.URL),
		str.FIELD_REQUEST_BODY: readRequestBodyForLogger(request, constants),
		str.FIELD_USER_AGENT:   request.UserAgent(),
	})
}
