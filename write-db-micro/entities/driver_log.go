package entities

import "time"

type DriverLog struct {
	DriverID  int       `bson:"driver_id"`
	Log       Log       `bson:"log"`
	CreatedAt time.Time `bson:"created_at"`
}

type Log struct {
	Lat float64 `bson:"lat"`
	Lon float64 `bson:"lon"`
}
