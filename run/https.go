package server

import (
	"context"
	"net/http"
)

func (s *Server) ServeHTTPS() {
	g := s.errG

	s.https = &http.Server{
		Addr:      s.cfg.HTTPSListener(),
		TLSConfig: s.certs.TLSConfig(),
		Handler:   tracing(nextRequestID)(logging(s.logger)(http.HandlerFunc(s.handleRedirect))),
	}

	g.Go(s.serveHTTPS())
}

func (s *Server) shutdownHTTPS(ctx context.Context) func() error {
	return func() error {
		s.log().Infof("Shutdown %s", s.https.Addr)

		return s.https.Shutdown(ctx)
	}
}

func (s *Server) serveHTTPS() func() error {
	return func() error {
		var err error

		s.log().Infof("Listening on %s", s.https.Addr)

		if err = s.https.ListenAndServeTLS("", ""); err == http.ErrServerClosed {
			return nil
		}

		return err
	}
}
