// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
	DefaultAcmeUrl = "https://acme-staging.api.letsencrypt.org/directory"

	// DefaultAcmeSkipTLS is skipping validation of acme directory
	DefaultAcmeSkipTLS = false

	// DefaultCacheSize
	DefaultCacheSize = 1000

	// DefaultCacheTTL (1min)
	DefaultCacheTTL = int64((5 * time.Second) / time.Nanosecond)
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
