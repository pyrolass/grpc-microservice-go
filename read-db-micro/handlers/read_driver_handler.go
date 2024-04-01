package handlers

import (
	"context"
	"fmt"

	types "github.com/pyrolass/grpc-microservice-go/proto"
)

type ReadGRPCDriverHandler struct {
	types.UnimplementedDriverReadServiceServer
}

func NewReadGRPCDriverHandler() *ReadGRPCDriverHandler {
	return &ReadGRPCDriverHandler{}
}

func (rs *ReadGRPCDriverHandler) GetDriverLogsDB(ctx context.Context, req *types.GetDriverLogRequest) (*types.GetDriverLogResponse, error) {

	fmt.Println("wawa wea")
	return nil, nil
}
