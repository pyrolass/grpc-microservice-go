package store

import (
	"context"
	"fmt"
	"time"

	types "github.com/pyrolass/grpc-microservice-go/proto"
	"github.com/pyrolass/grpc-microservice-go/read-db-micro/common"
	"github.com/pyrolass/grpc-microservice-go/read-db-micro/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type DriverStoreInterface interface {
	GetDriverLog(ctx context.Context, id int) (*types.GetDriverLogResponse, error)
}

type DriverStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewDriverStore(
	client *mongo.Client,
	coll *mongo.Collection,
) DriverStoreInterface {
	return &DriverStore{
		client: client,
		coll:   coll,
	}
}

func (d *DriverStore) GetDriverLog(ctx context.Context, id int) (*types.GetDriverLogResponse, error) {

	var driverLogs []entities.DriverLog

	cur, err := d.coll.Find(ctx, map[string]any{"driver_id": id, "created_at": map[string]any{"$gte": time.Now().Add(-time.Hour)}})

	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &driverLogs); err != nil {
		return nil, err
	}

	var distance float64

	for i := 0; i < len(driverLogs); i++ {

		if i == len(driverLogs)-1 {
			break
		}

		distance += common.Haversine(driverLogs[i].Log.Lat, driverLogs[i].Log.Lon, driverLogs[i+1].Log.Lat, driverLogs[i+1].Log.Lon)

	}

	var distanceDriverTook *types.GetDriverLogResponse = &types.GetDriverLogResponse{
		DriverId:          int32(id),
		DistanceTravelled: float32(distance),
	}

	fmt.Println(distanceDriverTook)

	return distanceDriverTook, nil
}
