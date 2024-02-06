package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/sboon-gg/svctl/internal/api"
	"github.com/sboon-gg/svctl/internal/daemon"
	"github.com/sboon-gg/svctl/svctl"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const (
	daemonPort = "50051"
)

type daemonOpts struct {
}

func newDaemonOpts() *daemonOpts {
	return &daemonOpts{}
}

func daemonCmd() *cobra.Command {
	opts := newDaemonOpts()

	cmd := &cobra.Command{
		Use:  "daemon",
		RunE: opts.Run,
	}

	return cmd
}

func (o *daemonOpts) Run(cmd *cobra.Command, args []string) error {
	d, err := daemon.Recover()
	if err != nil {
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", daemonPort))
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", daemonPort, err)
	}

	s := grpc.NewServer()
	svctl.RegisterServersServer(s, api.NewDaemonServer(d))
	log.Printf("gRPC server listening at %v", lis.Addr())

	go func() {
		for {
			select {
			case <-cmd.Context().Done():
				s.GracefulStop()
				return
			}
		}
	}()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(daemonCmd())
}
