package config

import (
	"github.com/sirupsen/logrus"
)

// InitLogging inits the logging for the app
func InitLogging() *logrus.Logger {
	logger := logrus.StandardLogger()
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})
	logger.SetLevel(logrus.DebugLevel)
	return logger
}
