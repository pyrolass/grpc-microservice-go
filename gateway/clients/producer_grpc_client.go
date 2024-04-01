package clients

import (
	types "github.com/pyrolass/grpc-microservice-go/proto"
	"google.golang.org/grpc"
)

type ProducerGRPCClient struct {
	Client types.DriverMessageServiceClient
	conn   *grpc.ClientConn
}

func NewProducerGRPCClient() *ProducerGRPCClient {
	// driver routes grpc
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := types.NewDriverMessageServiceClient(conn)

	return &ProducerGRPCClient{
		Client: client,
		conn:   conn,
	}
}

func (pc *ProducerGRPCClient) CloseConnection() {
	pc.conn.Close()
}
