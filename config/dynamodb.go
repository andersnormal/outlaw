package config

type DynamoDB struct {
	// Enable DynamoDB
	Enable bool

	// DynamoDB table
	Table string

	// DynamoDB endpoint
	Endpoint string

	// DynamoDBRegion
	Region string

	// DynamoDBAccessKey is the AWS Access Key
	AccessKey string

	// DynamoDBSecretKey is the AWS Secret Key
	SecretKey string
}
