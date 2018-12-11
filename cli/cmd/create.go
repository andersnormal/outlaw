package cmd

import (
	"context"
	"log"
	"time"

	"github.com/andersnormal/outlaw/cli/config"
	"github.com/andersnormal/outlaw/cli/dialer"
	pb "github.com/andersnormal/outlaw/proto"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const (
	defaultDomainType = pb.Domain_FULL_MATCH
)

var (
	enableWildcard bool
	path           string
	url            string
	params         bool
	statusCode     int32
)

func init() {
	CreateDomain.Flags().BoolVar(&enableWildcard, "wildcard", false, "wildcard")
	CreateDomain.Flags().BoolVar(&params, "parameters", false, "pass parameters")
	CreateDomain.Flags().StringVar(&path, "path", "*", "path match")
	CreateDomain.Flags().StringVar(&url, "url", "", "redirect url")
	CreateDomain.Flags().Int32Var(&statusCode, "status-code", 301, "http status code")
}

var CreateDomain = &cobra.Command{
	Use:   "create-domain",
	Short: "Creates new domain",
	Run: func(cmd *cobra.Command, args []string) {
		var opts []grpc.CallOption
		var err error

		if len(args) == 0 {
			return
		}

		d := &pb.Domain{
			Name: args[0],
			Type: defaultDomainType,
		}

		if enableWildcard {
			d.Type = pb.Domain_WILDCARD
		}

		r := &pb.Redirect{}

		if url != "" {
			m := &pb.Redirect_Match{
				Path: path,
				Type: pb.Redirect_Match_FULL_MATCH,
			}

			t := &pb.Redirect_Target{
				Url:        url,
				Parameters: params,
				Code:       pb.HTTPStatusCode(statusCode),
			}

			r.Match = m
			r.Target = t
		}

		d.Redirects = []*pb.Redirect{r}

		dial, err := dialer.NewDialer(config.C)
		if err != nil {
			log.Fatal(err)
		}
		defer dial.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if len(args) == 0 {
			log.Fatal(ErrNoDomain)
		}

		client := pb.NewOutlawClient(dial)
		in := &pb.CreateDomainRequest{Domain: d}

		resp, err := client.CreateDomain(ctx, in, opts...)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%s", resp.Domain.GetUuid())
	},
}
