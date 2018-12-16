package server

import (
	"context"
	"net"
	"time"

	pb "github.com/andersnormal/outlaw/proto"
	"github.com/andersnormal/outlaw/provider"

	// "github.com/gin-gonic/gin"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func (s *Server) ServeAPI() {
	g := s.errG

	s.api = grpc.NewServer()
	pb.RegisterOutlawServer(s.api, &API{s.config(), s.log()})

	g.Go(s.serveAPI())
}

func (s *Server) shutdownAPI(ctx context.Context) func() error {
	return func() error {
		s.api.GracefulStop()

		return nil
	}
}

func (s *Server) serveAPI() func() error {
	return func() error {
		var err error

		lis, err := net.Listen("tcp", s.cfg.APIListener())
		if err != nil {
			s.log().Error(err)
			return err
		}

		s.log().Infof("Listening on %s", lis.Addr())

		if err = s.api.Serve(lis); err != nil {
			return err
		}

		return nil
	}
}

func (a *API) log() *log.Entry {
	return a.logger
}

func (a *API) CreateDomain(ctx context.Context, req *pb.CreateDomainRequest) (*pb.CreateDomainResponse, error) {
	domain := req.Domain
	domain.Uuid = uuid.NewV5(uuid.NamespaceDNS, domain.Name).String()
	domain.Created = &timestamp.Timestamp{Seconds: time.Now().Unix()}
	domain.Modified = domain.Created

	domain, err := a.provider().CreateDomain(ctx, domain)
	if err != nil {
		a.log().Error(err)
		return &pb.CreateDomainResponse{}, ErrCreateDomainFailed.Err()
	}

	return &pb.CreateDomainResponse{Domain: domain}, nil
}

func (a *API) DeleteDomain(ctx context.Context, req *pb.DeleteDomainRequest) (*pb.DeleteDomainResponse, error) {
	domain := req.Domain

	if _, err := a.provider().DeleteDomain(ctx, domain); err != nil {
		a.log().Error(err)
		return &pb.DeleteDomainResponse{}, ErrDeleteDomainFailed.Err()
	}

	return &pb.DeleteDomainResponse{}, nil
}

func (a *API) CreateRedirect(ctx context.Context, req *pb.CreateRedirectRequest) (*pb.CreateRedirectResponse, error) {
	return &pb.CreateRedirectResponse{}, nil
}

func (a *API) DeleteRedirect(ctx context.Context, req *pb.DeleteDomainRequest) (*pb.DeleteDomainResponse, error) {
	return &pb.DeleteDomainResponse{}, nil
}

func (a *API) ListDomains(ctx context.Context, req *pb.ListDomainsRequest) (*pb.ListDomainsResponse, error) {
	domains, err := a.provider().ListDomains(ctx)
	if err != nil {
		a.log().Error(err)
		return &pb.ListDomainsResponse{}, ErrListDomainsFailed.Err()
	}

	return &pb.ListDomainsResponse{Domains: domains}, nil
}

func (a *API) GetDomain(ctx context.Context, req *pb.GetDomainRequest) (*pb.GetDomainResponse, error) {
	domain, err := a.provider().GetDomain(ctx, req.Domain)
	if err != nil {
		a.log().Error(err)
		return &pb.GetDomainResponse{}, ErrListDomainsFailed.Err()
	}

	return &pb.GetDomainResponse{Domain: domain}, nil
}

func (a *API) provider() provider.Provider {
	return a.cfg.Provider
}
