package cmd

import (
	"github.com/andersnormal/outlaw/version"
	"github.com/spf13/cobra"
)

func addSubCommands(root *cobra.Command) {
	// enable the bootstrap sub command
	root.AddCommand(bootstrapCmd)

	// enable version sub command
	root.AddCommand(version.Cmd)
}
