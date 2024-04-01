package clients

import (
	types "github.com/pyrolass/grpc-microservice-go/proto"
	"google.golang.org/grpc"
)

type DistanceGRPCClient struct {
	Client types.DriverQueryServiceClient
	conn   *grpc.ClientConn
}

func NewDistanceGRPCClient() *DistanceGRPCClient {
	// driver routes grpc
	conn, err := grpc.Dial("localhost:3004", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := types.NewDriverQueryServiceClient(conn)

	return &DistanceGRPCClient{
		Client: client,
	}
}

func (dc *DistanceGRPCClient) CloseConnection() {
	dc.conn.Close()
}
