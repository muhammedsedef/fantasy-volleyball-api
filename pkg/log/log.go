package log

import (
	"reflect"
	"strings"
	"time"

	"os"

	"github.com/sirupsen/logrus"
)

const lowestLogLevelPossible = logrus.DebugLevel

var logLevels = make(map[string]logrus.Level)

func SetLogLevel(loggerName string, newLevel string) {
	newLevel = strings.ToLower(newLevel)

	if newLevel == "debug" {
		logLevels[loggerName] = logrus.DebugLevel
	} else if newLevel == "info" {
		logLevels[loggerName] = logrus.InfoLevel
	} else if newLevel == "warn" {
		logLevels[loggerName] = logrus.WarnLevel
	} else if newLevel == "error" {
		logLevels[loggerName] = logrus.ErrorLevel
	} else {
		logLevels[loggerName] = logrus.InfoLevel
	}
}

func NewLoggerByType(loggerType reflect.Type, defaultFields map[string]interface{}) Logger {
	return NewLoggerByName(loggerType.String(), defaultFields)
}

func NewLoggerByName(loggerName string, defaultFields map[string]interface{}) Logger {
	logger := new(LogrusLogger)

	logger.name = loggerName
	logger.defaultFields = getDefaultFields(defaultFields)

	logger.underlyingLogger = logrus.StandardLogger()

	if suppressLogs() {
		logger.underlyingLogger.Level = logrus.FatalLevel
	} else {
		logger.underlyingLogger.Level = lowestLogLevelPossible
	}

	if isBuildMode() {
		logrus.SetFormatter(&PrettyJSONFormatter{TimestampFormat: time.RFC3339Nano})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: time.RFC3339Nano})
	}

	addLogLevelIfMissing(loggerName, "info")

	return logger
}

func getDefaultFields(defaultFields map[string]interface{}) map[string]interface{} {
	fields := make(map[string]interface{})

	for field, value := range defaultFields {
		if !isSensitiveField(field) {
			fields[field] = value
		}
	}
	return fields
}

func mapFieldsToEntry(entry *logrus.Entry, fields map[string]interface{}) {
	for field, value := range fields {
		if !isSensitiveField(field) {
			escapedField := strings.Replace(field, ".", "_", -1)
			entry.Data[escapedField] = value
		}
	}
}

func isSensitiveField(field string) bool {
	return field == "Authorization"
}

func addLogLevelIfMissing(loggerName string, newLevel string) {
	if logLevels[loggerName] == logrus.PanicLevel {
		SetLogLevel(loggerName, newLevel)
	}
}

func isDebugEnabled(loggerName string) bool {
	return logLevels[loggerName] >= logrus.DebugLevel
}

func isInfoEnabled(loggerName string) bool {
	return logLevels[loggerName] >= logrus.InfoLevel
}

func isWarnEnabled(loggerName string) bool {
	return logLevels[loggerName] >= logrus.WarnLevel
}

func isErrorEnabled(loggerName string) bool {
	return logLevels[loggerName] >= logrus.ErrorLevel
}

func isBuildMode() bool {
	return os.Getenv("BUILD_MODE") == "1"
}

func suppressLogs() bool {
	return os.Getenv("SUPPRESS_LOGS") == "1"
}
