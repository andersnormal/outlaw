package server

import (
	"net"
	"net/http"
	"net/url"

	pb "github.com/andersnormal/outlaw/proto"

	log "github.com/sirupsen/logrus"
)

// handler returns the redirector function for the HTTP server
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
		match, target := redirect.GetMatch(), redirect.GetTarget()
		if match == nil {
			break
		}

		// break here
		if match.GetType() == pb.Redirect_Match_WILDCARD {
			s.redirect(w, r, match, target)

			return
		}

		// try to match path
		if r.URL.Path == match.Path || match.Path == "*" {
			s.redirect(w, r, match, target)

			return
		}
	}

	http.NotFound(w, r)
}

func (s *Server) redirect(w http.ResponseWriter, r *http.Request, match *pb.Redirect_Match, target *pb.Redirect_Target) {
	log := s.log().WithFields(log.Fields{
		"path":   match.GetPath(),
		"target": target.GetUrl(),
	})

	u, err := url.Parse(target.GetUrl())
	if err != nil {
		http.Error(w, "Bad redirect", http.StatusBadRequest)
	}

	// if the request parameters should be passed
	if target.GetParameters() {
		u.RawQuery = r.URL.RawQuery
	}

	log.Info("Redirecting")

	// really do the redirect
	http.Redirect(w, r, u.String(), int(target.GetCode()))
}
