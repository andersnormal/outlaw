package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Version string

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "Every software has its version and this is outlaw's",
	RunE:  runE,
}

func runE(cmd *cobra.Command, args []string) error {
	var err error

	fmt.Printf("v%s", Version)

	return err
}
