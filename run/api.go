package server

import (
	"context"
	// "encoding/json"
	// "fmt"
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

	// nextRequestID := func() string {
	// 	return fmt.Sprintf("%d", time.Now().UnixNano())
	// }

	// gin.SetMode(gin.ReleaseMode)

	// // register api router
	// router := gin.Default()
	// router.GET("/health", s.health)
	// router.GET("/export", s.exportDomains)
	// router.POST("/import", s.importDomains)
	// router.GET("/domain", s.fetchAllDomains)
	// router.GET("/domain/:name", s.fetchDomain)
	// router.POST("/domain", s.registerDomain)
	// router.DELETE("/domain/:name", s.purgeDomain)

	// lis, err := net.Listen("tcp", s.cfg.APIListener())
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

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

// health handler
// func (s *Server) health(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"message": "ok"})
// }

// // exportDomains exports the domains
// func (s *Server) exportDomains(c *gin.Context) {
// 	domains, err := s.cfg.Provider.FetchAll()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error while fetching domains"})
// 		return
// 	}

// 	export := &provider.ExportDomains{
// 		Domains: domains,
// 	}

// 	c.JSON(http.StatusOK, export)
// }

// // importDomains imports a domain export set
// func (s *Server) importDomains(c *gin.Context) {
// 	if c.Request.Body == nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Please send a request body"})
// 		return
// 	}

// 	var export provider.ExportDomains

// 	if err := json.NewDecoder(c.Request.Body).Decode(&export); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
// 		return
// 	}
// 	defer c.Request.Body.Close()

// 	if err := s.cfg.Provider.DeleteAllDomains(); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database operation failed"})
// 		return
// 	}

// 	if err := s.cfg.Provider.Import(&export); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database operation failed"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{})
// }

// // purgeDomain deletes a domain entry
// func (s *Server) purgeDomain(w http.ResponseWriter, r *http.Request, ps gin.Params) {
// 	name := ps.ByName("name")
// 	domain, err := s.cfg.Provider.FetchByDomain(name)

// 	if domain == nil || err != nil {
// 		sendJSONMessage(w, "not found", 404)
// 		return
// 	}

// 	if _, err = s.cfg.Provider.DeleteByDomain(name); err != nil {
// 		// log.Error(err)
// 		sendJSONMessage(w, "Error while deleting domain", 500)
// 		return
// 	}

// 	// s.cfg.Provider.DeleteTLSCacheEntry(name)

// 	sendJSONMessage(w, "ok", 204)
// }

// // fetchAllDomains return a list of all domains
// func (s *Server) fetchAllDomains(w http.ResponseWriter, r *http.Request, _ gin.Params) {
// 	domains, err := s.cfg.Provider.FetchAll()

// 	if err != nil {
// 		sendJSONMessage(w, "Error while fetching domains", 500)
// 		return
// 	}

// 	sendJSON(w, domains, 200)
// }

// func (s *Server) fetchDomain(w http.ResponseWriter, r *http.Request, ps gin.Params) {
// 	name := ps.ByName("name")
// 	domain, err := s.cfg.Provider.FetchByDomain(name)

// 	if err != nil {
// 		sendJSONMessage(w, "not found", 404)
// 		return
// 	}

// 	sendJSON(w, domain, 200)
// }

// func (s *Server) registerDomain(w http.ResponseWriter, r *http.Request, _ gin.Params) {
// 	if r.Body == nil {
// 		sendJSONMessage(w, "Please send a request body", 400)
// 		return
// 	}

// 	var domain provider.Domain

// 	if err := json.NewDecoder(r.Body).Decode(&domain); err != nil {
// 		sendJSONMessage(w, "Invalid request body", 400)
// 		return
// 	}

// 	domain.ID = uuid.Must(uuid.NewV4(), provider.ErrInvalidID).String()
// 	domain.Created = time.Now().Format(time.RFC3339)
// 	domain.Modified = domain.Created

// 	// validate
// 	if errList := domain.Validate(); len(errList) > 0 {
// 		errMsg := ""
// 		for _, err := range errList {
// 			errMsg = errMsg + err.Error() + ". "
// 		}
// 		sendJSONMessage(w, errMsg, 400)
// 		return
// 	}

// 	// insert new domain
// 	// if err := s.cfg.Provider.InsertDomain(domain); err != nil {
// 	// 	// log.Error(err)
// 	// 	sendJSONMessage(w, "Can't store document", 500)
// 	// 	return
// 	// }

// 	sendJSONMessage(w, "ok", 201)
// }
