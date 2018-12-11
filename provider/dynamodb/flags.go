package dynamodb

import (
	"github.com/andersnormal/outlaw/config"
	"github.com/spf13/cobra"
)

func AddFlags(cmd *cobra.Command, cfg *config.Config) {
	// enable DynamoDB
	cmd.PersistentFlags().BoolVar(&cfg.DynamoDB.Enable, "dynamodb", false, "enable dynamodb")

	// DynamoDB table
	cmd.PersistentFlags().StringVar(&cfg.DynamoDB.Table, "dynamodb-table", "Domains", "DynamoDB Table")

	// DynamoDB Access Key
	cmd.PersistentFlags().StringVar(&cfg.DynamoDB.AccessKey, "dynamodb-access-key", "", "DynamoDB Access Key")

	// DynamoDB Secret Key
	cmd.PersistentFlags().StringVar(&cfg.DynamoDB.SecretKey, "dynamodb-secret-key", "", "DynamoDB Secret Key")

	// DynamoDB Region
	cmd.PersistentFlags().StringVar(&cfg.DynamoDB.Region, "dynamodb-region", "eu-west-1", "DynamoDB Region")

	// DynamoDB Endpoint
	cmd.PersistentFlags().StringVar(&cfg.DynamoDB.Endpoint, "dynamodb-endpoint", "", "DynamoDB Endpoint")
}
