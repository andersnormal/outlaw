package mongo

import (
	"github.com/andersnormal/outlaw/config"

	"github.com/andersnormal/lru"
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

// Mongo
type Mongo struct {
	// config to use with the server
	cfg     *config.Config
	logger  *log.Entry
	session *mgo.Session

	// a simple LRU cache
	lru lru.Cache
}

type Cachable struct {
	Key  string
	Data []byte
}
