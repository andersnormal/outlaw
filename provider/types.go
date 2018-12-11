package provider

import (
	"context"

	pb "github.com/andersnormal/outlaw/proto"

	"golang.org/x/crypto/acme/autocert"
)

// Provider defines the interface to a certificate provider
type Provider interface {
	Bootstrap(ctx context.Context) error

	AllowHostPolicy(ctx context.Context, host string) error

	CreateDomain(ctx context.Context, domain *pb.Domain) (*pb.Domain, error)
	DeleteDomain(ctx context.Context, domain *pb.Domain) (bool, error)
	UpdateDomain(ctx context.Context, domain *pb.Domain) (*pb.Domain, error)
	ListDomains(ctx context.Context) ([]*pb.Domain, error)
	GetDomain(ctx context.Context, domain *pb.Domain) (*pb.Domain, error)

	Name() string

	// confirm to the cache interface
	autocert.Cache
}

// DomainList db entry
// type DomainList struct {
// 	Domains []Domain `json:"domains"`
// }

// PathMappingEntry model
type PathMappingEntry struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// PathList model
type PathList []PathMappingEntry

// Domain db entry
// type Domain struct {
// 	ID           string    `json:"id"`
// 	Name         string    `json:"domain"`
// 	PathMapping  *PathList `json:"path"`
// 	Redirect     string    `json:"redirect"`
// 	Promotable   bool      `json:"promotable"`
// 	Wildcard     bool      `json:"wildcard"`
// 	Certificate  string    `json:"certificate"`
// 	RedirectCode int       `json:"code"`
// 	Description  string    `json:"description"`
// 	Created      string    `json:"created"`
// 	Modified     string    `json:"modified"`
// }

// ExportDomains model
// type ExportDomains struct {
// 	Domains []Domain `json:"domains"`
// }
