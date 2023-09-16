package logger

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

var logger *logrus.Logger

func NewLogger() {
	logLevel := logrus.DebugLevel
	log := logrus.New()
	log.SetLevel(logLevel)
	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename: fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02")),
		MaxSize:  50, // MB
		MaxAge:   28, // DAYS
		Level:    logLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat:   "2006-01-02 15:04:05",
			PrettyPrint:       true,
			DisableHTMLEscape: true,
		},
	})

	if err != nil {
		logrus.Panic(err)
	}

	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
		PrettyPrint:       true,
		DisableHTMLEscape: true,
	})

	log.AddHook(rotateFileHook)

	logger = log
}

func WriteLog(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}

func GetLogger() *logrus.Logger {
	return logger
}
