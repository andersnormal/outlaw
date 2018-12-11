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
	"syscall"
	"time"

	"github.com/andersnormal/outlaw/provider"
	log "github.com/sirupsen/logrus"
)

// Config contains a configuration for Voskhod
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
