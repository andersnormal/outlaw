package certs

import (
	"crypto/tls"
	"net/http"

	"github.com/andersnormal/outlaw/cache"
)

// Manager is an interface to manage certificates
type Manager interface {
	// GetCertificate is retrieving a certificate from the manager
	GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error)
	// ServeHTTP is serving the autocert handler
	ServeHTTP(fallback http.Handler, w http.ResponseWriter, r *http.Request)
	// TLSConfig exposes the tls config
	TLSConfig() *tls.Config
	// Cache
	Cache() cache.Cache
}
