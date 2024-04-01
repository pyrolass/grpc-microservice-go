package entities

type DriverEntity struct {
	DriverId int     `json:"driver_id" validate:"required"`
	Lat      float64 `json:"lat" validate:"required"`
	Lon      float64 `json:"lon" validate:"required"`
}

type DriverLog struct {
	DriverId int
	Distance float64
}
