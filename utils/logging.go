package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

func SetupLogging() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}
