// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
