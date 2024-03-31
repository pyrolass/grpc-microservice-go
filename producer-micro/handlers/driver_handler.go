package handlers

import (
	"context"

	types "github.com/pyrolass/grpc-microservice-go/proto"
	"github.com/sirupsen/logrus"
)

type GRPCDriverHandler struct {
	types.UnimplementedDriverServiceServer
}

func NewGRPCDriverHandler() *GRPCDriverHandler {
	return &GRPCDriverHandler{}
}

func (d *GRPCDriverHandler) InsertDriverLog(ctx context.Context, req *types.InsertDriverLogRequest) (*types.InsertDriverLogResponse, error) {
	logrus.Infof("InsertDriverLog called with %v", req)
	return &types.InsertDriverLogResponse{
		Success: true,
	}, nil
}
