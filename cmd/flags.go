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

package cmd

import (
	"github.com/andersnormal/outlaw/config"
	"github.com/spf13/cobra"
)

func addFlags(cmd *cobra.Command, cfg *config.Config) {
	// enable verbose output
	cmd.PersistentFlags().BoolVar(&cfg.Verbose, "verbose", config.DefaultVerbose, "enable verbose output")

	// timeout for client operations
	cmd.Flags().DurationVar(&cfg.Timeout, "timeout", config.DefaultTimeout, "timeout")

	// enable staging
	cmd.Flags().StringVar(&cfg.AcmeUrl, "acme-url", config.DefaultAcmeUrl, "ACME directory url")

	// skip acme tls verification
	cmd.Flags().BoolVar(&cfg.AcmeSkipTLS, "acme-skip-tls", config.DefaultAcmeSkipTLS, "ACME skip tls verification")

	// API Port
	cmd.Flags().IntVar(&cfg.APIPort, "api-port", config.DefaultAPIPort, "api port")

	// HTTP Port
	cmd.Flags().IntVar(&cfg.HTTPPort, "http-port", config.DefaultHTTPPort, "http port")

	// HTTPS Port
	cmd.Flags().IntVar(&cfg.HTTPSPort, "https-port", config.DefaultHTTPSPort, "https port")

	// Host to listen on
	cmd.Flags().StringVar(&cfg.Host, "host", config.DefaultHost, "host")

	// Cache size to use
	cmd.Flags().IntVar(&cfg.CacheSize, "cache", config.DefaultCacheSize, "cache size")
}
