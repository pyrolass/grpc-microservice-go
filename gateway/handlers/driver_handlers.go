package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pyrolass/grpc-microservice-go/gateway/entities"
	types "github.com/pyrolass/grpc-microservice-go/proto"
)

type DriverHandlerInterface interface {
	InsertDriverLog(c *fiber.Ctx) error
}

type DriverHandler struct {
	gClient types.DriverServiceClient
}

func NewDriverHandler(gClient types.DriverServiceClient) DriverHandlerInterface {
	return &DriverHandler{
		gClient: gClient,
	}
}

func (d *DriverHandler) InsertDriverLog(c *fiber.Ctx) error {
	var driver entities.DriverEntity
	if err := c.BodyParser(&driver); err != nil {
		return err
	}

	res, err := d.gClient.InsertDriverLog(c.Context(), &types.InsertDriverLogRequest{
		DriverId: int32(driver.DriverId),
		Lat:      driver.Lat,
		Lon:      driver.Lon,
	})

	if err != nil {
		return err
	}

	if !res.Success {
		return c.Status(500).JSON(map[string]string{
			"msg": "failed to insert driver log",
		})
	}

	return c.JSON(map[string]string{
		"message": "destination received",
	})
}
