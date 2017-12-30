package crypto

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.Out = os.Stdout
}

func EnableDebugLogs(enable bool) {
	level := logrus.InfoLevel

	if enable {
		level = logrus.DebugLevel
	}

	log.SetLevel(level)
}

func Log() *logrus.Logger {
	return log
}
