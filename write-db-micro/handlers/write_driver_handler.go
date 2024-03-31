package handlers

import (
	"context"

	types "github.com/pyrolass/grpc-microservice-go/proto"

	"github.com/sirupsen/logrus"
)

type WriteGRPCDriverHandler struct {
	types.UnimplementedDriverMessageServiceServer
}

func NewWriteGRPCDriverHandler() *WriteGRPCDriverHandler {
	return &WriteGRPCDriverHandler{}
}

func (d *WriteGRPCDriverHandler) InsertDriverLog(ctx context.Context, req *types.InsertDriverLogRequest) (*types.InsertDriverLogResponse, error) {
	logrus.Infof("InsertDriverLog called with %v", req)

	return &types.InsertDriverLogResponse{
		Success: true,
	}, nil
}
