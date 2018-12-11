package cache

import (
	"context"

	pb "github.com/andersnormal/outlaw/proto"
	"github.com/andersnormal/outlaw/provider"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/acme/autocert"
)

// Cache is representing a cache
type Cache interface {
	AllowHostPolicy(ctx context.Context, host string) error
	GetDomain(ctx context.Context, domain string) (*pb.Domain, error)

	autocert.Cache
}

// cache represents a cache instance
type cache struct {
	p provider.Provider
	l *log.Entry
}
