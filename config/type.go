package config

import (
	"syscall"
	"time"

	"github.com/andersnormal/outlaw/provider"
	log "github.com/sirupsen/logrus"
)

// Config contains a configuration for outlaw
type Config struct {
	// Verbose toggles the verbosity
	Verbose bool

	// LogLevel is the level with with to log for this config
	LogLevel log.Level

	// ReloadSignal
	ReloadSignal syscall.Signal

	// TermSignal
	TermSignal syscall.Signal

	// KillSignal
	KillSignal syscall.Signal

	// Timeout of the runtime
	Timeout time.Duration

	// Bootstrap
	Bootstrap bool

	// DomainsTableName is the name of the database table
	DomainsTableName string

	// LogFormat
	LogFormat string

	// HTTPPort is the port for HTTP
	HTTPPort int

	// HTTPSPort is the port for HTTPS
	HTTPSPort int

	// APIPort is the port for API
	APIPort int

	// CacheSize is the size of the cache
	CacheSize int

	// CacheTTL is the timeout for the cache
	CacheTTL int64

	// Host is the host to listen on
	Host string

	// Provider connection
	Provider provider.Provider

	// Acme
	Acme bool

	// AcmeStaging
	AcmeUrl string

	// AcmeSkipTLS
	AcmeSkipTLS bool

	// DynamoDB
	DynamoDB DynamoDB

	// Mongo
	Mongo Mongo
}
