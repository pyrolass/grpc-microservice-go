package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pyrolass/grpc-microservice-go/gateway/entities"
)

type DriverHandlerInterface interface {
	InsertDriverLog(c *fiber.Ctx) error
}

type DriverHandler struct {
}

func NewDriverHandler() DriverHandlerInterface {
	return &DriverHandler{}
}

func (d *DriverHandler) InsertDriverLog(c *fiber.Ctx) error {
	var driver entities.DriverEntity
	if err := c.BodyParser(&driver); err != nil {
		return err
	}

	return c.JSON(map[string]string{
		"message": "destination received",
	})
}
