package server

import (
	"context"
	"crypto/tls"
	"net/http"
	"sync"

	"github.com/andersnormal/outlaw/certs"
	"github.com/andersnormal/outlaw/config"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

// Listener describes the interface to a server
type Listener interface {
	// ServeHTTP is starting the HTTP listener
	ServeHTTP()
	// ServeHTTPS is starting the HTTPS listener
	ServeHTTPS()
	// ServeAPI is starting the API listener
	ServeAPI()
	// Wait is waiting for everything to end :-)
	Wait() error
}

type listener struct {
	// config to use with the server
	cfg *config.Config

	// logger attached to server
	logger *log.Entry

	// manager for the certs
	certs certs.Manager

	// error Group
	errG *errgroup.Group

	// error Context
	errCtx context.Context

	// http
	http *http.Server

	// https
	https *http.Server

	// api
	api *grpc.Server

	// tls config
	tls *tls.Config

	// lock is used to safely access the client
	lock sync.RWMutex
}

// Server represents a listener
type Server listener

// API
type API struct {
	// config to be used with the api
	cfg *config.Config

	// logger
	logger *log.Entry
}
