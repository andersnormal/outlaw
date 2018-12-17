package config

import (
	"fmt"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	// DefaultDomainsTableName is the default name for the domains table.
	DefaultDomainsTableName = "Domains"

	// DefaultLogFormat is the default format of the logging.
	// The default is to log to JSON.
	DefaultLogFormat = "json"

	// DefaultLogLevel is the default logging level.
	DefaultLogLevel = log.WarnLevel

	// DefaultTermSignal is the signal to term the agent.
	DefaultTermSignal = syscall.SIGTERM

	// DefaultReloadSignal is the default signal for reload.
	DefaultReloadSignal = syscall.SIGHUP

	// DefaultKillSignal is the default signal for termination.
	DefaultKillSignal = syscall.SIGINT

	// DefaultVerbose is the default verbosity.
	DefaultVerbose = false

	// DefaultTimeout is the default time to configure the runtime
	DefaultTimeout = 1 * time.Minute

	// DefaultHTTPSPort is the default port for HTTPS
	DefaultHTTPSPort = 443

	// DefaultHTTPPort is the default port for HTTP
	DefaultHTTPPort = 80

	// DefaultAPIPort is the default port for API
	DefaultAPIPort = 8888

	// DefaultHost to listen on
	DefaultHost = ""

	// DefaultAcme
	DefaultAcme = true

	// DefaultAcmeUrl (default is actually the staging, you should try here)
	DefaultAcmeUrl = "https://acme-v01.api.letsencrypt.org/directory"

	// DefaultAcmeSkipTLS is skipping validation of acme directory
	DefaultAcmeSkipTLS = false

	// DefaultCacheSize
	DefaultCacheSize = 1000

	// DefaultCacheTTL (1min)
	DefaultCacheTTL = int64((5 * time.Second) / time.Nanosecond)

	// DefaultBootstrap
	DefaultBootstrap = false
)

// New returns a new Config
func New() *Config {
	return &Config{
		Verbose:          DefaultVerbose,
		LogLevel:         DefaultLogLevel,
		ReloadSignal:     DefaultReloadSignal,
		TermSignal:       DefaultTermSignal,
		KillSignal:       DefaultKillSignal,
		Timeout:          DefaultTimeout,
		DomainsTableName: DefaultDomainsTableName,
		LogFormat:        DefaultLogFormat,
		Host:             DefaultHost,
		APIPort:          DefaultAPIPort,
		HTTPPort:         DefaultHTTPPort,
		HTTPSPort:        DefaultHTTPSPort,
		Acme:             DefaultAcme,
		AcmeUrl:          DefaultAcmeUrl,
		AcmeSkipTLS:      DefaultAcmeSkipTLS,
		CacheSize:        DefaultCacheSize,
		CacheTTL:         DefaultCacheTTL,
		Bootstrap:        DefaultBootstrap,
	}
}

// HTTPListener returns the listener for HTTP
func (c *Config) HTTPListener() string {
	return fmt.Sprintf("%s:%d", c.Host, c.HTTPPort)
}

// HTTPSListener returns the listener for HTTPS
func (c *Config) HTTPSListener() string {
	return fmt.Sprintf("%s:%d", c.Host, c.HTTPSPort)
}

// APIListener returns the listener for API
func (c *Config) APIListener() string {
	return fmt.Sprintf("%s:%d", c.Host, c.APIPort)
}
