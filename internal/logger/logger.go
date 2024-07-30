package logger

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Level = logrus.DebugLevel
}

func Debug(msg ...string) {
	log.Debug(msg)
}

func Info(msg ...string) {
	log.Info(msg)
}

func Warning(msg ...string) {
	log.Warning(msg)
}

func Error(msg ...string) {
	log.Error(msg)
}

func Fatal(err ...error) {
	log.Fatal(err)
}
