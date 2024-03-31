package handlers

import (
	"context"

	"github.com/pyrolass/grpc-microservice-go/producer-micro/producer"
	types "github.com/pyrolass/grpc-microservice-go/proto"

	"github.com/sirupsen/logrus"
)

type ProducerGRPCDriverHandler struct {
	types.UnimplementedDriverMessageServiceServer
	p producer.ProducerInterface
}

func NewProducerGRPCDriverHandler(p producer.ProducerInterface) *ProducerGRPCDriverHandler {
	return &ProducerGRPCDriverHandler{
		p: p,
	}
}

func (d *ProducerGRPCDriverHandler) InsertDriverLog(ctx context.Context, req *types.InsertDriverLogRequest) (*types.InsertDriverLogResponse, error) {
	logrus.Infof("InsertDriverLog called with %v", req)

	err := d.p.ProduceDriverDistance(req)

	if err != nil {
		logrus.Errorf("Error producing driver distance: %v", err)
	}

	return &types.InsertDriverLogResponse{
		Success: true,
	}, nil
}
