/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/sboon-gg/svctl/svctl"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type registerOpts struct {
	*serverOpts
}

func newRegisterOpts() *registerOpts {
	return &registerOpts{
		serverOpts: newServerOpts(),
	}
}

func registerCmd() *cobra.Command {
	opts := newRegisterOpts()

	cmd := &cobra.Command{
		Use:   "register",
		Short: "register a new user",
		RunE:  opts.Run,
	}

	opts.AddFlags(cmd)

	return cmd
}

func (o *registerOpts) AddFlags(cmd *cobra.Command) {
	o.serverOpts.AddFlags(cmd)
}

func (o *registerOpts) Run(cmd *cobra.Command, args []string) error {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()
	c := svctl.NewServersClient(conn)

	ctx, cancel := context.WithTimeout(cmd.Context(), time.Second)
	defer cancel()

	path, err := o.Path()
	if err != nil {
		return err
	}

	r, err := c.Register(ctx, &svctl.ServerOpts{Path: path})
	if err != nil {
		return fmt.Errorf("error calling function Start: %v", err)
	}

	cmd.Printf("Server status: %v\n", r.GetStatus().String())
	return nil
}

func init() {
	rootCmd.AddCommand(registerCmd())
}
