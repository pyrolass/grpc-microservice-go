package handlers

import (
	"context"

	types "github.com/pyrolass/grpc-microservice-go/proto"
	"github.com/pyrolass/grpc-microservice-go/read-db-micro/store"
)

type ReadGRPCDriverHandler struct {
	types.UnimplementedDriverReadServiceServer
	store store.DriverStoreInterface
}

func NewReadGRPCDriverHandler(store store.DriverStoreInterface) *ReadGRPCDriverHandler {
	return &ReadGRPCDriverHandler{
		store: store,
	}
}

func (rs *ReadGRPCDriverHandler) GetDriverLogsDB(ctx context.Context, req *types.GetDriverLogRequest) (*types.GetDriverLogResponse, error) {
	data, err := rs.store.GetDriverLog(ctx, int(req.DriverId))

	if err != nil {
		return nil, err
	}

	return data, nil
}
