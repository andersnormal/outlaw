package acme

import (
	"crypto/tls"
	"net/http"

	"github.com/andersnormal/outlaw/cache"
	"github.com/andersnormal/outlaw/certs"
	"github.com/andersnormal/outlaw/config"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/acme"
	"golang.org/x/crypto/acme/autocert"
)

var _ certs.Manager = (*AcmeManager)(nil)

// NewManager creates a new instance
func NewManager(cfg *config.Config, c cache.Cache) certs.Manager {
	// acme client
	client := &acme.Client{
		DirectoryURL: cfg.AcmeUrl,
	}

	// dev, disable checking the acme tls certificate
	if cfg.AcmeSkipTLS {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		ht := &http.Client{Transport: tr}

		client.HTTPClient = ht
	}

	// create manager
	m := &AcmeManager{
		logger: log.WithFields(log.Fields{
			"url": cfg.AcmeUrl,
		}),
	}

	// autocert manager
	autocert := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: c.AllowHostPolicy,
		Client:     client,
		Cache:      c,
	}
	m.manager = autocert
	m.cache = c

	return m
}

func (a *AcmeManager) Cache() cache.Cache {
	return a.cache
}

// GetCertificate wrapper for the cert getter
func (a *AcmeManager) GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return a.manager.GetCertificate(hello)
}

// ServeHTTP is serving the Acme challange responses
func (a *AcmeManager) ServeHTTP(fallback http.Handler, w http.ResponseWriter, r *http.Request) {
	a.manager.HTTPHandler(fallback).ServeHTTP(w, r)
}

// TLSConfig
func (a *AcmeManager) TLSConfig() *tls.Config {
	// use autocert config
	c := a.manager.TLSConfig()

	// Pass in a cert manager if you want one set
	// this will only be used if the server Certificates are empty
	c.GetCertificate = a.GetCertificate

	// VersionTLS11 or VersionTLS12 would exclude many browsers
	// inc. Android 4.x, IE 10, Opera 12.17, Safari 6
	// So unfortunately not acceptable as a default yet
	// Current default here for clarity
	c.MinVersion = tls.VersionTLS10

	// // Causes servers to use Go's default ciphersuite preferences,
	// // which are tuned to avoid attacks. Does nothing on clients.
	c.PreferServerCipherSuites = true

	// // Only use curves which have assembly implementations
	c.CurvePreferences = []tls.CurveID{
		tls.CurveP256,
		tls.X25519, // Go 1.8 only
	}

	return c
}
