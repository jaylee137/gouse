package log

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/constants"
	"github.com/thuongtruong1009/gouse/strings"
)

func detectError(err interface{}) string {
	switch e := err.(type) {
	case error:
		return e.Error()
	case string:
		return e
	default:
		return fmt.Sprintf("%v", err)
	}
}

func Write(level, msg, filePath string, err ...interface{}) {
	logParam := &Param{
		Prefix:  level,
		Message: "",
		Output:  filePath,
	}
	logId := strings.RandomStr(8)

	if len(err) > 0 {
		errStr := detectError(err)
		logParam.Message = fmt.Sprintf("Message: %s - ID: %s - Error: %s - Date: \n", msg, logId, errStr)
	} else {
		logParam.Message = fmt.Sprintf("Message: %s - ID: %s - Date: \n", msg, logId)
	}

	Handle(logParam.Output, fmt.Sprintf("%s %s", logParam.Prefix, logParam.Message), logParam.Output)
}

func WriteError(msg string, err interface{}) {
	Write(constants.ERROR_LOG_PREFIX, msg, constants.ERROR_LOG_PATH, err)
}

func WriteInfo(msg string) {
	Write(constants.INFO_LOG_PREFIX, msg, constants.INFO_LOG_PATH)
}

func WriteAccess(msg string) {
	Write(constants.ACCESS_LOG_PREFIX, msg, constants.ACCESS_LOG_PATH)
}

func WriteCustom(prefix, msg string, filePath ...string) {
	Write(prefix, msg, filePath[0])
}
