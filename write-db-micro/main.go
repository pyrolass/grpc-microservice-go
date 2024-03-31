package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	types "github.com/pyrolass/grpc-microservice-go/proto"
	"github.com/pyrolass/grpc-microservice-go/write-db-micro/handlers"
	"github.com/pyrolass/grpc-microservice-go/write-db-micro/store"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

func main() {
	grpcListenAddr := flag.String("grpc-addr", ":3003", "server listen address")

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")

	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	logrus.Info("Connected to MongoDB!")

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	driverStore := store.NewDriverStore(
		client,
		client.Database("driver").Collection("logs"),
	)

	err = makeGRPCTransport(*grpcListenAddr, driverStore)

	if err != nil {
		logrus.Fatalf("Error creating GRPC transport: %v", err)
	}
}

func makeGRPCTransport(listenAddr string, store store.DriverStoreInterface) error {
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

	driverHandler := handlers.NewWriteGRPCDriverHandler(store)

	types.RegisterDriverWriteServiceServer(grpcServer, driverHandler)

	return grpcServer.Serve(ln)

}
