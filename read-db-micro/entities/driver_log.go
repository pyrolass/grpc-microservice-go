package entities

type DriverLog struct {
	DriverID int `bson:"driver_id"`
	Log      Log `bson:"log"`
}

type Log struct {
	Lat float64 `bson:"lat"`
	Lon float64 `bson:"lon"`
}
