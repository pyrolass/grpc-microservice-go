package handlers

import (
	"context"
	"fmt"

	"github.com/pyrolass/grpc-microservice-go/distance-micro/clients"
	types "github.com/pyrolass/grpc-microservice-go/proto"
)

type DistanceGRPCDriverHandler struct {
	types.UnimplementedDriverQueryServiceServer
	readClient *clients.ReadGRPCClient
}

func NewDistanceGRPCDriverHandler(readClient *clients.ReadGRPCClient) *DistanceGRPCDriverHandler {
	return &DistanceGRPCDriverHandler{
		readClient: readClient,
	}
}

func (h *DistanceGRPCDriverHandler) GetDriverLogs(ctx context.Context, req *types.GetDriverLogRequest) (*types.GetDriverLogResponse, error) {
	fmt.Println("test")
	res, err := h.readClient.Client.GetDriverLogsDB(ctx, req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
