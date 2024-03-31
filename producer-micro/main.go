package main

import (
	"flag"
	"net"

	"github.com/pyrolass/grpc-microservice-go/producer-micro/handlers"
	types "github.com/pyrolass/grpc-microservice-go/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	grpcListenAddr := flag.String("grpc-addr", ":3001", "server listen address")

	err := makeGRPCTransport(*grpcListenAddr)

	if err != nil {
		logrus.Fatalf("Error creating GRPC transport: %v", err)
	}

}

func makeGRPCTransport(listenAddr string) error {
	// make the tcp
	logrus.Infof("GRPC transport starting on %s", listenAddr)

	ln, err := net.Listen("tcp", listenAddr)

	if err != nil {
		logrus.Fatalf("Failed to listen: %v", err)
		return err
	}

	defer ln.Close()
	// make grpc server

	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)

	driverHandler := handlers.NewGRPCDriverHandler()

	types.RegisterDriverServiceServer(grpcServer, driverHandler)

	return grpcServer.Serve(ln)

}
