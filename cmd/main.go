package cmd

import (
	"flag"
	"fmt"
	"github.com/balinux8/movo/core/hello"
	"github.com/balinux8/movo/pkg/services"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

var serverCmd = &cobra.Command{
	Use: "server",
	Short: "Start movie service",
	Long: "Start movie service",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Start server here")
		if err := run(); err != nil {
			fmt.Println(err)
		}
	},
}

func init()  {
	RootCmd.AddCommand(serverCmd)
}

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost9090", "endpoint of YourService")
)

func run() error  {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := hello.RegisterHelloHandlerFromEndpoint(ctx, mux, ":8082", opts)
	if err != nil {
		return err
	}

	closing := make(chan struct{})
	go func() {
		if err := startGrpcServer(":8082"); err != nil {
			fmt.Println(err)
		}

		closing <- struct{}{}
	}()

	go func() {
		if err := http.ListenAndServe(":8081", mux); err != nil {
			fmt.Println(err)
		}

		closing <- struct{}{}
	}()

	go func() {
		if err := http.ListenAndServe(":8081", mux); err != nil {
			fmt.Println(err)
		}
		closing <- struct{}{}
	}()

	<- closing

	return nil
}

func startGrpcServer(address string) error  {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	fmt.Printf(`Starting GRPC server at: "%s"`, address)

	server := grpc.NewServer()
	hello.RegisterHelloServer(server, services.HelloServiceImpl{})
	err = server.Serve(listen)

	return err
}