package cmd

import (
	"context"
	"os"

	"github.com/andersnormal/outlaw/cache"
	"github.com/andersnormal/outlaw/certs"
	"github.com/andersnormal/outlaw/certs/acme"
	server "github.com/andersnormal/outlaw/run"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func runE(cmd *cobra.Command, args []string) error {
	var err error
	var root = new(root)
	var m certs.Manager

	// init logger
	root.logger = log.WithFields(log.Fields{})

	// create sys channel
	root.sys = make(chan os.Signal, 1)
	root.exit = make(chan int, 1)

	// create root context
	root.ctx, root.cancel = context.WithCancel(context.Background())
	defer root.cancel()

	// watch syscalls and cancel upon need
	go root.watchSignals(cfg)

	// create cache
	c := cache.New(cfg.Provider)

	// only support acme
	switch {
	default:
		m = acme.NewManager(cfg, c)
	}

	s := server.NewServer(root.ctx, cfg, m)
	// start https
	s.ServeHTTPS()
	// start http
	s.ServeHTTP()
	// start api
	s.ServeAPI()

	// await error, or shutdown on signal
	err = s.Wait()

	// noop
	return err
}
