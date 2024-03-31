package main

import (
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var topic = "driver_distance"

func main() {
	cons, err := NewKafkaConsumer(topic)

	go func() {
		cons.Start()
	}()

	if err != nil {
		logrus.Fatalf("Error creating Kafka consumer: %v", err)
	}

	err = makeGRPCTransport(":3002")

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

	// driverHandler := handlers.NewGRPCDriverHandler(p)

	// types.RegisterDriverServiceServer(grpcServer, driverHandler)

	return grpcServer.Serve(ln)

}
