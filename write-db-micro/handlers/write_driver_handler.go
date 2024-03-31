package handlers

import (
	"context"

	types "github.com/pyrolass/grpc-microservice-go/proto"

	"github.com/sirupsen/logrus"
)

type WriteGRPCDriverHandler struct {
	types.UnimplementedDriverWriteServiceServer
}

func NewWriteGRPCDriverHandler() *WriteGRPCDriverHandler {
	return &WriteGRPCDriverHandler{}
}

func (d *WriteGRPCDriverHandler) InsertDriverLogDB(ctx context.Context, req *types.InsertLogListRequest) (*types.InsertDriverLogResponse, error) {
	logrus.Infof("InsertDriverLog called with %v", req)

	return &types.InsertDriverLogResponse{
		Success: true,
	}, nil
}
