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

	// Bootstrap the database upon launch
	cmd.Flags().BoolVar(&cfg.Bootstrap, "bootstrap", config.DefaultBootstrap, "bootstrap database upon start")
}
