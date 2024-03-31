package store

import (
	"context"

	"github.com/pyrolass/grpc-microservice-go/write-db-micro/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type DriverStoreInterface interface {
	InsertDriverLog(ctx context.Context, data []entities.DriverLog) error
}

type DriverStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewDriverStore(
	client *mongo.Client,
	coll *mongo.Collection,
) *DriverStore {
	return &DriverStore{
		client: client,
		coll:   coll,
	}
}

func (d *DriverStore) InsertDriverLog(ctx context.Context, data []entities.DriverLog) error {

	interfaceSlice := make([]interface{}, len(data))

	for i, v := range data {
		interfaceSlice[i] = v
	}

	_, err := d.coll.InsertMany(ctx, interfaceSlice)

	if err != nil {
		return err
	}

	return nil
}
