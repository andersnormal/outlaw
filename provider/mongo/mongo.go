package mongo

import (
	"context"
	"time"

	"github.com/andersnormal/outlaw/config"
	pb "github.com/andersnormal/outlaw/proto"
	"github.com/andersnormal/outlaw/provider"

	"github.com/andersnormal/lru"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

const (
	defaultDomainCollection = "domains"
	defaultCacheCollection  = "cache"
	defaultName             = "MongoDB"
)

var _ provider.Provider = (*Mongo)(nil)

// NewMongo creates a new instance of the MongoDB provider
func NewMongo(cfg *config.Config, l lru.Cache) (provider.Provider, error) {
	var err error

	dial := &mgo.DialInfo{
		Addrs:    []string{cfg.Mongo.Endpoint},
		Timeout:  60 * time.Second,
		Database: cfg.Mongo.Database,
	}

	if cfg.Mongo.Password != "" && cfg.Mongo.Username != "" {
		dial.Username = cfg.Mongo.Username
		dial.Password = cfg.Mongo.Password
	}

	if cfg.Mongo.AuthDatabase != "" {
		dial.Database = cfg.Mongo.AuthDatabase
	}

	session, err := mgo.DialWithInfo(dial)
	if err != nil {
		return nil, err
	}

	// create logger
	logger := log.WithFields(log.Fields{
		"mongo.database": cfg.Mongo.Database,
		"mongo.endpoint": cfg.Mongo.Endpoint,
	})

	return &Mongo{cfg, logger, session, l}, err
}

// Name returns the name of the provider
func (d *Mongo) Name() string {
	return defaultName
}

// prepareTable checks for the main table
func (d *Mongo) Bootstrap(ctx context.Context) error {
	cols := map[string]*mgo.CollectionInfo{
		defaultDomainCollection: &mgo.CollectionInfo{},
		defaultCacheCollection:  &mgo.CollectionInfo{},
	}

	for c, info := range cols {
		if err := d.db().C(c).Create(info); err != nil {
			return err
		}
	}

	return nil
}

// CreateDomain creates a domain in DynamoDB.
func (d *Mongo) CreateDomain(ctx context.Context, domain *pb.Domain) (*pb.Domain, error) {
	var err error

	// insert domain
	err = d.db().C(defaultDomainCollection).Insert(domain)
	if err != nil {
		return nil, err
	}

	return domain, err
}

func (d *Mongo) DeleteDomain(ctx context.Context, domain *pb.Domain) (bool, error) {
	var err error

	err = d.db().C(defaultDomainCollection).Remove(bson.M{"name": domain.GetName()})
	if err != nil && err != mgo.ErrNotFound {
		return false, err
	}

	err = d.Delete(ctx, domain.GetName())
	if err != nil && err != mgo.ErrNotFound {
		return false, err
	}

	return true, nil
}

func (d *Mongo) UpdateDomain(ctx context.Context, domain *pb.Domain) (*pb.Domain, error) {
	return nil, nil
}

func (d *Mongo) GetDomain(ctx context.Context, domain *pb.Domain) (*pb.Domain, error) {
	var err error
	var res *pb.Domain

	q := d.db().C(defaultDomainCollection).Find(bson.M{"name": domain.GetName()})
	err = q.One(&res)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (d *Mongo) ListDomains(ctx context.Context) ([]*pb.Domain, error) {
	var err error
	var domains []*pb.Domain

	q := d.db().C(defaultDomainCollection).Find(struct{}{})
	err = q.All(&domains)
	if err != nil {
		return nil, err
	}

	return domains, err
}

func (d *Mongo) AllowHostPolicy(ctx context.Context, host string) error {
	_, err := d.GetDomain(ctx, &pb.Domain{Name: host})

	return err
}

func (d *Mongo) Get(ctx context.Context, key string) ([]byte, error) {
	var (
		err  error
		item *Cachable
	)

	q := d.db().C(defaultCacheCollection).Find(bson.M{"key": key})
	err = q.One(&item)
	if err != nil {
		return nil, err
	}

	return item.Data, err
}

func (d *Mongo) Put(ctx context.Context, key string, data []byte) error {
	var err error

	item := &Cachable{key, data}

	// insert to cache or update
	_, err = d.db().C(defaultCacheCollection).Upsert(bson.M{"key": key}, item)
	if err != nil {
		return err
	}

	return nil
}

func (d *Mongo) Delete(ctx context.Context, key string) error {
	return d.db().C(defaultCacheCollection).Remove(bson.M{"key": key})
}

func (d *Mongo) db() *mgo.Database {
	return d.session.DB(d.cfg.Mongo.Database)
}
