package server

import (
	"net"
	"net/http"

	pb "github.com/andersnormal/outlaw/proto"
)

func (s *Server) handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.certs.ServeHTTP(http.HandlerFunc(s.handleRedirect), w, r)
	})
}

// handle normal redirect request on http
func (s *Server) handleRedirect(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	domain, err := s.certs.Cache().GetDomain(r.Context(), host)
	if domain == nil || err != nil {
		http.NotFound(w, r)

		return
	}

	if domain.Type == pb.Domain_FULL_MATCH && host != domain.Name {
		http.NotFound(w, r)

		return
	}

	for _, redirect := range domain.Redirects {
		target := redirect.Target

		http.Redirect(w, r, target.Url, int(target.Code))

		return
	}

	http.NotFound(w, r)
}
