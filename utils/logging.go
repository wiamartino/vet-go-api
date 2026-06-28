package utils

import (
	"os"
	"strings"

	"go-vet/config"

	"github.com/sirupsen/logrus"
)

// SetupLogging configures logrus using values from config.AppConfig.
// Must be called after config.LoadConfig().
func SetupLogging() {
	logrus.SetOutput(os.Stdout)

	// Use config if available, otherwise fall back to defaults
	cfg := config.AppConfig
	if cfg == nil {
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
		logrus.SetLevel(logrus.InfoLevel)
		return
	}

	// Configure log format
	switch strings.ToLower(cfg.Logging.Format) {
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	default:
		logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	}

	// Configure log level
	level, err := logrus.ParseLevel(cfg.Logging.Level)
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel)
		logrus.Warningf("Invalid LOG_LEVEL %q, defaulting to info", cfg.Logging.Level)
	} else {
		logrus.SetLevel(level)
	}
}
