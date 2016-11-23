package util

import (
	"github.com/Sirupsen/logrus"
)

// Log util

// Logger is the default logger with info level
var Logger = logrus.New()

// Short name use in util package
var log = Logger.WithFields(logrus.Fields{
	"pkg": "x.util",
})

func init() {
	Logger.Formatter = &logrus.TextFormatter{ForceColors: true}
	Logger.Level = logrus.InfoLevel
}

// UseVerboseLog set logger level to debug
func UseVerboseLog() {
	Logger.Level = logrus.DebugLevel
	log.Debug("enable debug logging")
}
