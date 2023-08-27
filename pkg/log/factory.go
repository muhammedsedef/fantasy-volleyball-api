package log

import (
	configuration "fantasy-volleyball-api/appconfig"
	"reflect"
)

func GetLogger(loggerType reflect.Type) Logger {
	loggerMetaData := make(map[string]interface{})
	loggerMetaData["app_name"] = "fantasy-volleyball-api"
	loggerMetaData["facility"] = "fantasy-volleyball-api"
	loggerMetaData["teamName"] = "folksdev"
	loggerMetaData["profile"] = configuration.Env

	logger := NewLoggerByType(loggerType, loggerMetaData)
	return logger
}
