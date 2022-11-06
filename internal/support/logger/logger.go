package logger

import (
	"fmt"
	"gin_tonic/internal/support/localContext"
	"os"
	"path/filepath"
	"time"
)

const (
	EMERGENCY = "EMERGENCY"
	ALERT     = "ALERT"
	CRITICAL  = "CRITICAL"
	ERROR     = "ERROR"
	WARNING   = "WARNING"
	NOTICE    = "NOTICE"
	INFO      = "INFO"
	DEBUG     = "DEBUG"
)

func EmergencyLog(context localContext.LocalContext, header string, message string) {
	log(context, header, message, EMERGENCY)
}

func AlertLog(context localContext.LocalContext, header string, message string) {
	log(context, header, message, ALERT)
}

func CriticalLog(context localContext.LocalContext, header string, message string) {
	log(context, header, message, CRITICAL)
}

func ErrorLog(context localContext.LocalContext, header string, message string) {
	log(context, header, message, ERROR)
}

func WarningLog(context localContext.LocalContext, header string, message string) {
	log(context, header, message, WARNING)
}

func NoticeLog(context localContext.LocalContext, header string, message string) {
	log(context, header, message, NOTICE)
}

func InfoLog(context localContext.LocalContext, header string, message string) {
	log(context, header, message, INFO)
}

func DebugLog(context localContext.LocalContext, header string, message string) {
	log(context, header, message, DEBUG)
}

func log(context localContext.LocalContext, header string, message string, level string) {
	logFile := findOrCreateLogFile(context)
	defer logFile.Close()

	timeNow := time.Now()
	_, err := logFile.WriteString(
		"[" + timeNow.Format(`01-02-2006 15:04:05`) + "] [level: " + level +
			"] [header: " + header + "]\n" + message + "\n",
	)
	context.InternalServerError(err)
}

func findOrCreateLogFile(context localContext.LocalContext) *os.File {
	var logFile *os.File
	var err error

	year, month, day := time.Now().Date()
	logFilePath := fmt.Sprintf("storage/logs/date-%v-%s-%v.log", year, month, day)
	files, _ := filepath.Glob(logFilePath)
	isFind := len(files) > 0

	if isFind {
		logFile, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_WRONLY, 0666)
	} else {
		logFile, err = os.Create(logFilePath)
	}
	context.InternalServerError(err)

	return logFile
}
