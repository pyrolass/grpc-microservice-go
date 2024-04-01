package clients

import (
	types "github.com/pyrolass/grpc-microservice-go/proto"
	"google.golang.org/grpc"
)

type ReadGRPCClient struct {
	Client types.DriverReadServiceClient
	conn   *grpc.ClientConn
}

func NewReadGRPCClient() *ReadGRPCClient {
	// driver routes grpc
	conn, err := grpc.Dial("localhost:3005", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := types.NewDriverReadServiceClient(conn)

	return &ReadGRPCClient{
		Client: client,
	}
}

func (dc *ReadGRPCClient) CloseConnection() {
	dc.conn.Close()
}
