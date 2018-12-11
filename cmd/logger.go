package cmd

import (
	"github.com/andersnormal/outlaw/config"
	log "github.com/sirupsen/logrus"
)

// setupLog prepares the logger instance
func setupLog(cfg *config.Config) {
	switch cfg.LogFormat {
	case "text":
		log.SetFormatter(&log.TextFormatter{})
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.SetFormatter(&log.JSONFormatter{})
	}

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Only log the warning severity or above.
	log.SetLevel(cfg.LogLevel)

	// Set the format of the logger

	// if we should output verbose
	if cfg.Verbose {
		log.SetLevel(log.InfoLevel)
	}
}
