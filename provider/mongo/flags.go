package mongo

import (
	"github.com/andersnormal/outlaw/config"

	"github.com/spf13/cobra"
)

func AddFlags(cmd *cobra.Command, cfg *config.Config) {
	// enable DynamoDB
	cmd.PersistentFlags().BoolVar(&cfg.Mongo.Enable, "mongo", false, "enable mongo")

	// Database
	cmd.PersistentFlags().StringVar(&cfg.Mongo.Database, "mongo-database", "outlaw", "Mongo database")

	// Endpoint
	cmd.PersistentFlags().StringVar(&cfg.Mongo.Endpoint, "mongo-endpoint", "localhost", "Mongo endpoint")

	// Username
	cmd.PersistentFlags().StringVar(&cfg.Mongo.Username, "mongo-username", "", "Mongo username")

	// Password
	cmd.PersistentFlags().StringVar(&cfg.Mongo.Password, "mongo-password", "", "Mongo password")

	// Auth database
	cmd.PersistentFlags().StringVar(&cfg.Mongo.AuthDatabase, "mongo-auth-database", "", "Mongo auth database")
}
