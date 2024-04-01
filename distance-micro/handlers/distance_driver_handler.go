package handlers

import (
	"context"

	types "github.com/pyrolass/grpc-microservice-go/proto"
)

type DistanceGRPCDriverHandler struct {
	types.UnimplementedDriverQueryServiceServer
}

func NewDistanceGRPCDriverHandler() *DistanceGRPCDriverHandler {
	return &DistanceGRPCDriverHandler{}
}

func (h *DistanceGRPCDriverHandler) GetDriverLogs(ctx context.Context, req *types.GetDriverLogRequest) (*types.None, error) {
	return &types.None{}, nil
}
