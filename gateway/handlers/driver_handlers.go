package handlers

import "github.com/gofiber/fiber/v2"

type DriverHandlerInterface interface {
	InsertDriverLog(c *fiber.Ctx) error
}

type DriverHandler struct {
	DriverId string  `json:"driver_id"`
	Lat      float64 `json:"lat"`
	Lon      float64 `json:"lon"`
}

func NewDriverHandler() DriverHandlerInterface {
	return &DriverHandler{}
}

func (d *DriverHandler) InsertDriverLog(c *fiber.Ctx) error {
	return nil
}
