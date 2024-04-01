package handlers

import (
	"context"
	"time"

	types "github.com/pyrolass/grpc-microservice-go/proto"
	"github.com/pyrolass/grpc-microservice-go/write-db-micro/entities"
	"github.com/pyrolass/grpc-microservice-go/write-db-micro/store"

	"github.com/sirupsen/logrus"
)

type WriteGRPCDriverHandler struct {
	types.UnimplementedDriverWriteServiceServer
	store store.DriverStoreInterface
}

func NewWriteGRPCDriverHandler(store store.DriverStoreInterface) *WriteGRPCDriverHandler {
	return &WriteGRPCDriverHandler{
		store: store,
	}
}

func (d *WriteGRPCDriverHandler) InsertDriverLogDB(ctx context.Context, req *types.InsertLogListRequest) (*types.InsertDriverLogResponse, error) {
	logrus.Infof("InsertDriverLog called with %v", req)

	var logs []entities.DriverLog

	for _, log := range req.Logs {
		logs = append(logs, entities.DriverLog{
			DriverID:  int(log.DriverId),
			CreatedAt: time.Now(),
			Log: entities.Log{
				Lat: log.Lat,
				Lon: log.Lon,
			},
		})
	}

	err := d.store.InsertDriverLog(ctx, logs)

	if err != nil {
		return &types.InsertDriverLogResponse{
			Success: false,
		}, err
	}

	return &types.InsertDriverLogResponse{
		Success: true,
	}, nil
}
