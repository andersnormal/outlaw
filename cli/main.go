package main

import (
	"math/rand"
	"time"

	"github.com/andersnormal/outlaw/cli/cmd"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

var rootCmd = &cobra.Command{
	Use: "outlaw",
	Run: func(cmd *cobra.Command, args []string) {
		return
	},
}

func init() {
	rand.Seed(time.Now().UnixNano())

	// silence on the root cmd
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	rootCmd.AddCommand(cmd.CreateDomain)
	rootCmd.AddCommand(cmd.DeleteDomain)
	rootCmd.AddCommand(cmd.ListDomains)
	rootCmd.AddCommand(cmd.GetDomain)
}

func main() {
	rootCmd.Execute()
}
