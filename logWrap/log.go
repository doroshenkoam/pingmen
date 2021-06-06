package logWrap

import (
	"github.com/sirupsen/logrus"
)

// SetLogLevel - set log level, default warning level
func SetLogLevel(lvl string) {
	level, err := logrus.ParseLevel(lvl)
	if err != nil {
		level = logrus.WarnLevel
	}

	logrus.SetLevel(level)
}

// SetBaseFields - set base fields to log
func SetBaseFields(p, f string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"package":  p,
		"function": f,
	})
}
