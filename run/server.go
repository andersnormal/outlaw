package server

import (
	"context"
	"time"

	"github.com/andersnormal/outlaw/certs"
	"github.com/andersnormal/outlaw/config"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var _ Listener = (*Server)(nil)

func NewServer(ctx context.Context, cfg *config.Config, manager certs.Manager) Listener {
	g, gtx := errgroup.WithContext(ctx)

	s := &Server{
		cfg:    cfg,
		errCtx: gtx,
		errG:   g,
		certs:  manager,
		logger: log.WithFields(log.Fields{
			"host": cfg.Host,
		}),
	}

	return s
}

func (s *Server) Wait() error {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
		case <-s.errCtx.Done():
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			g, gtx := errgroup.WithContext(ctx)

			if s.https != nil {
				g.Go(s.shutdownHTTPS(gtx))
			}

			if s.http != nil {
				g.Go(s.shutdownHTTP(gtx))
			}

			if s.api != nil {
				g.Go(s.shutdownAPI(gtx))
			}

			return g.Wait()
		}
	}
}

func (s *Server) config() *config.Config {
	return s.cfg
}

func (s *Server) log() *log.Entry {
	return s.logger
}
