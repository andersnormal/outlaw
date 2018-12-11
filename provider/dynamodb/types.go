package dynamodb

import (
	"github.com/andersnormal/outlaw/config"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	log "github.com/sirupsen/logrus"
)

// DynamoDB model
type DynamoDB struct {
	// config to use with the server
	cfg    *config.Config
	logger *log.Entry
	client dynamodbiface.DynamoDBAPI
}

// DynamoConnection model
type DynamoConnection struct {
	Endpoint  string
	Key       string
	Secret    string
	TableName string
	Region    string
}

// DomainList db entry
type DomainList struct {
	Domains []Domain `json:"domains"`
}

// PathMappingEntry model
type PathMappingEntry struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// PathList model
type PathList []PathMappingEntry

// Domain db entry
type Domain struct {
	ID           string    `json:"id"`
	Name         string    `json:"domain"`
	PathMapping  *PathList `json:"path"`
	Redirect     string    `json:"redirect"`
	Promotable   bool      `json:"promotable"`
	Wildcard     bool      `json:"wildcard"`
	Certificate  string    `json:"certificate"`
	RedirectCode int       `json:"code"`
	Description  string    `json:"description"`
	Created      string    `json:"created"`
	Modified     string    `json:"modified"`
}

// ExportDomains model
type ExportDomains struct {
	Domains []Domain `json:"domains"`
}
