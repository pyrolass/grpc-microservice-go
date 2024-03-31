package main

import (
	"flag"
	"net"

	"github.com/pyrolass/grpc-microservice-go/producer-micro/handlers"
	"github.com/pyrolass/grpc-microservice-go/producer-micro/producer"
	types "github.com/pyrolass/grpc-microservice-go/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var topic = "driver_distance"

func main() {
	grpcListenAddr := flag.String("grpc-addr", ":3001", "server listen address")

	p, err := producer.NewKafkaProducer(topic)

	if err != nil {
		logrus.Fatalf("Error creating Kafka producer: %v", err)
	}

	err = makeGRPCTransport(*grpcListenAddr, p)

	if err != nil {
		logrus.Fatalf("Error creating GRPC transport: %v", err)
	}

}

func makeGRPCTransport(listenAddr string, p producer.ProducerInterface) error {
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

	driverHandler := handlers.NewGRPCDriverHandler(p)

	types.RegisterDriverServiceServer(grpcServer, driverHandler)

	return grpcServer.Serve(ln)

}
