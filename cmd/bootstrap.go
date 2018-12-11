package cmd

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use:     "bootstrap",
	Short:   "Bootstrap the database",
	RunE:    bootstrapRunE,
	PreRunE: preRunE,
}

func bootstrapRunE(cmd *cobra.Command, args []string) error {
	// var err error
	var root = new(root)

	// init logger
	root.logger = log.WithFields(log.Fields{})

	// create root context
	root.ctx, root.cancel = context.WithCancel(context.Background())
	defer root.cancel()

	// call to bootstrap
	if err := cfg.Provider.Bootstrap(root.ctx); err != nil {
		root.logger.Fatal(err)
	}

	return nil // noop
}
