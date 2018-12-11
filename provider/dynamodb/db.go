package dynamodb

import (
	"context"

	"github.com/andersnormal/outlaw/config"
	pb "github.com/andersnormal/outlaw/proto"
	"github.com/andersnormal/outlaw/provider"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	dyn "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	log "github.com/sirupsen/logrus"
)

var _ provider.Provider = (*DynamoDB)(nil)

const (
	defaultName = "DynamoDB"
)

// NewDynamoDB creates a new instance
func NewDynamoDB(cfg *config.Config) (provider.Provider, error) {
	var err error
	var creds *credentials.Credentials

	// use env provider by default
	creds = credentials.NewEnvCredentials()

	// if there are options set to overwrite
	if cfg.DynamoDB.AccessKey != "" && cfg.DynamoDB.SecretKey != "" {
		creds = credentials.NewStaticCredentials(cfg.DynamoDB.AccessKey, cfg.DynamoDB.SecretKey, "")
	}

	config := &aws.Config{
		Credentials: creds,
		Region:      aws.String(cfg.DynamoDB.Region),
		Endpoint:    aws.String(cfg.DynamoDB.Endpoint),
	}

	sess, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}

	client := dyn.New(sess)
	logger := log.WithFields(log.Fields{
		"dyn.region":   cfg.DynamoDB.Region,
		"dyn.table":    cfg.DynamoDB.Table,
		"dyn.endpoint": cfg.DynamoDB.Endpoint,
	})

	return &DynamoDB{cfg, logger, client}, nil
}

func (d *DynamoDB) Name() string {
	return defaultName
}

// prepareTable checks for the main table
func (d *DynamoDB) Bootstrap(ctx context.Context) error {
	var err error

	d.logger.Info("Bootstraping DynamoDB")

	dbDomainTableCreate := &dyn.CreateTableInput{
		TableName: aws.String(d.cfg.DomainsTableName),
		KeySchema: []*dyn.KeySchemaElement{
			{AttributeName: aws.String("name"), KeyType: aws.String("HASH")},
		},
		AttributeDefinitions: []*dyn.AttributeDefinition{
			{AttributeName: aws.String("name"), AttributeType: aws.String("S")},
		},
		ProvisionedThroughput: &dyn.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}

	_, err = d.client.CreateTableWithContext(ctx, dbDomainTableCreate)
	if err != nil {
		d.logger.Error(err)
	}

	return err
}

// CreateDomain creates a domain in DynamoDB.
func (d *DynamoDB) CreateDomain(ctx context.Context, domain *pb.Domain) (*pb.Domain, error) {
	var err error

	mm, err := dynamodbattribute.MarshalMap(domain)
	if err != nil {
		return nil, err
	}

	_, err = d.client.PutItem(&dyn.PutItemInput{
		Item:      mm,
		TableName: aws.String(d.cfg.DynamoDB.Table),
	})

	return domain, err
}

func (d *DynamoDB) GetDomain(ctx context.Context, domain *pb.Domain) (*pb.Domain, error) {
	var err error

	return nil, err
}

func (d *DynamoDB) DeleteDomain(ctx context.Context, domain *pb.Domain) (bool, error) {
	var err error

	out, err := d.client.DeleteItem(&dyn.DeleteItemInput{
		Key: map[string]*dyn.AttributeValue{
			"name": {
				S: aws.String(domain.Name),
			},
		},
		TableName: aws.String(d.cfg.DynamoDB.Table),
	})

	return out != nil && err == nil, err
}

func (d *DynamoDB) UpdateDomain(ctx context.Context, domain *pb.Domain) (*pb.Domain, error) {
	return nil, nil
}

func (d *DynamoDB) ListDomains(ctx context.Context) ([]*pb.Domain, error) {
	var domains []*pb.Domain

	return domains, nil
}

func (d *DynamoDB) Get(ctx context.Context, key string) ([]byte, error) {
	return make([]byte, 0), nil
}

func (d *DynamoDB) Put(ctx context.Context, key string, data []byte) error {
	return nil
}

func (d *DynamoDB) Delete(ctx context.Context, key string) error {
	return nil
}

func (d *DynamoDB) AllowHostPolicy(ctx context.Context, host string) error {
	return nil
}

// UpdateCertificateData updates the cert data if a domain entry exist
// func (d *DynamoDB) UpdateCertificateData(domain string, data []byte) error {
// 	_, err := d.Service.UpdateItem(&dyn.UpdateItemInput{
// 		TableName: aws.String(dbDomainTableName),
// 		Key: map[string]*dyn.AttributeValue{
// 			"domain": {
// 				S: aws.String(domain),
// 			},
// 		},
// 		UpdateExpression: aws.String("set certificate = :c"),
// 		ReturnValues:     aws.String("UPDATED_NEW"),
// 		ExpressionAttributeValues: map[string]*dyn.AttributeValue{
// 			":c": {
// 				S: aws.String(string(data)),
// 			},
// 		},
// 	})

// 	return err
// }

// 	return out != nil && err == nil, err
// }

// FetchByDomain items from domains table
// func (d *DynamoDB) FetchByDomain(domain string) (*provider.Domain, error) {
// 	res, err := d.Service.GetItem(&dyn.GetItemInput{
// 		TableName: aws.String(dbDomainTableName),
// 		Key: map[string]*dyn.AttributeValue{
// 			"domain": {
// 				S: aws.String(domain),
// 			},
// 		},
// 	})

// 	if err != nil {
// 		return nil, fmt.Errorf("Error while getting item. %v", err)
// 	}

// 	domainRes := &provider.Domain{}
// 	if err = dynamodbattribute.UnmarshalMap(res.Item, &domainRes); err == nil {
// 		return domainRes, nil
// 	}

// 	return nil, nil
// }

// DeleteAllDomains deletes all items from the domains table
// func (d *DynamoDB) DeleteAllDomains() error {
// 	domains, err := d.FetchAll()

// 	if err != nil {
// 		return err
// 	}

// 	for _, do := range domains {
// 		_, err = d.DeleteByDomain(do.Name)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// // Import imports a export set
// func (d *DynamoDB) Import(e *provider.ExportDomains) error {
// 	for _, do := range e.Domains {
// 		mm, err := dynamodbattribute.MarshalMap(do)

// 		if err != nil {
// 			return err
// 		}

// 		_, err = d.Service.PutItem(&dyn.PutItemInput{
// 			Item:      mm,
// 			TableName: aws.String(dbDomainTableName),
// 		})
// 	}

// 	return nil
// }
