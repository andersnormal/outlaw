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

var DeleteDomain = &cobra.Command{
	Use:   "delete-domain",
	Short: "Deletes a domain",
	Run: func(cmd *cobra.Command, args []string) {
		var opts []grpc.CallOption
		var err error

		dial, err := dialer.NewDialer(config.C)
		if err != nil {
			log.Fatal(err)
		}
		defer dial.Close()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client := pb.NewOutlawClient(dial)
		in := &pb.DeleteDomainRequest{Domain: &pb.Domain{
			Name: args[0],
		}}

		resp, err := client.DeleteDomain(ctx, in, opts...)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%s", resp.Domain.GetUuid())
	},
}
