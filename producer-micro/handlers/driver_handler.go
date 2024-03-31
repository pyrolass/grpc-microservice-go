package handlers

import (
	"context"

	"github.com/pyrolass/grpc-microservice-go/producer-micro/producer"
	types "github.com/pyrolass/grpc-microservice-go/proto"

	"github.com/sirupsen/logrus"
)

type GRPCDriverHandler struct {
	types.UnimplementedDriverServiceServer
	p producer.ProducerInterface
}

func NewGRPCDriverHandler(p producer.ProducerInterface) *GRPCDriverHandler {
	return &GRPCDriverHandler{
		p: p,
	}
}

func (d *GRPCDriverHandler) InsertDriverLog(ctx context.Context, req *types.InsertDriverLogRequest) (*types.InsertDriverLogResponse, error) {
	logrus.Infof("InsertDriverLog called with %v", req)

	err := d.p.ProduceDriverDistance(req)

	if err != nil {
		logrus.Errorf("Error producing driver distance: %v", err)
	}

	return &types.InsertDriverLogResponse{
		Success: true,
	}, nil
}
