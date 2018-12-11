package cmd

import (
	"github.com/andersnormal/outlaw/provider/dynamodb"
	"github.com/andersnormal/outlaw/provider/mongo"

	"github.com/andersnormal/lru"
	"github.com/spf13/cobra"
)

func preRunE(cmd *cobra.Command, args []string) error {
	var err error

	// create the LRU cache
	l, err := lru.NewLRU(cfg.CacheSize)
	if err != nil {
		return err
	}

	// we switch context here in terms of enablement
	switch {
	case cfg.Mongo.Enable:
		cfg.Provider, err = mongo.NewMongo(cfg, l)
	case cfg.DynamoDB.Enable:
		cfg.Provider, err = dynamodb.NewDynamoDB(cfg)
	default:
		return ErrNoProvider
	}

	return err
}
