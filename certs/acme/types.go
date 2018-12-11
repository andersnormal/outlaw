package acme

import (
	"github.com/andersnormal/outlaw/cache"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/acme/autocert"
)

type AcmeManager struct {
	cache   cache.Cache
	manager *autocert.Manager
	logger  *log.Entry
}
