package logger

import (
	"fmt"
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

func EmergencyLog(header string, message string) error {
	return log(header, message, EMERGENCY)
}

func AlertLog(header string, message string) error {
	return log(header, message, ALERT)
}

func CriticalLog(header string, message string) error {
	return log(header, message, CRITICAL)
}

func ErrorLog(header string, message string) error {
	return log(header, message, ERROR)
}

func WarningLog(header string, message string) error {
	return log(header, message, WARNING)
}

func NoticeLog(header string, message string) error {
	return log(header, message, NOTICE)
}

func InfoLog(header string, message string) error {
	return log(header, message, INFO)
}

func DebugLog(header string, message string) error {
	return log(header, message, DEBUG)
}

func log(header string, message string, level string) error {
	logFile, err := findOrCreateLogFile()
	defer logFile.Close()
	if err != nil {
		return err
	}

	timeNow := time.Now()
	_, err = logFile.WriteString(
		"[" + timeNow.Format(`01-02-2006 15:04:05`) + "] [level: " + level +
			"] [header: " + header + "]\n" + message + "\n",
	)
	if err != nil {
		return err
	}

	return nil
}

func findOrCreateLogFile() (*os.File, error) {
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
	if err != nil {
		return nil, err
	}

	return logFile, err
}
