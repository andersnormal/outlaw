package cmd

import (
	"fmt"
	"os"

	"github.com/andersnormal/outlaw/config"
	// "github.com/andersnormal/outlaw/provider/dynamodb"
	"github.com/andersnormal/outlaw/provider/mongo"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfg *config.Config
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "outlaw",
	Short:   "The outlaw redirector.",
	Long:    ``,
	PreRunE: preRunE,
	RunE:    runE,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// init config
	cfg = config.New()

	// initialize cobra
	cobra.OnInitialize(initConfig)

	// add sub commands
	addSubCommands(RootCmd)

	// adding flags
	addFlags(RootCmd, cfg)

	// adding DynamoDB flags
	// dynamodb.AddFlags(RootCmd, cfg)

	// adding MongoDB flags
	mongo.AddFlags(RootCmd, cfg)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match

	// setup the logging behavior
	setupLog(cfg)
}
