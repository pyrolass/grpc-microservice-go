package handlers

import (
	"context"
	"fmt"

	types "github.com/pyrolass/grpc-microservice-go/proto"
)

type DistanceGRPCDriverHandler struct {
	types.UnimplementedDriverQueryServiceServer
}

func NewDistanceGRPCDriverHandler() *DistanceGRPCDriverHandler {
	return &DistanceGRPCDriverHandler{}
}

func (h *DistanceGRPCDriverHandler) GetDriverLogs(ctx context.Context, req *types.GetDriverLogRequest) (*types.GetDriverLogResponse, error) {
	fmt.Println("test")
	return nil, nil
}
