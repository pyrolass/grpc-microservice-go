package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pyrolass/grpc-microservice-go/gateway/entities"
	types "github.com/pyrolass/grpc-microservice-go/proto"
)

type DriverHandlerInterface interface {
	InsertDriverLog(c *fiber.Ctx) error
	GetDriverLogs(c *fiber.Ctx) error
}

type DriverHandler struct {
	pc types.DriverMessageServiceClient
	dc types.DriverQueryServiceClient
}

func NewDriverHandler(pc types.DriverMessageServiceClient, dc types.DriverQueryServiceClient) DriverHandlerInterface {
	return &DriverHandler{
		pc: pc,
		dc: dc,
	}
}

func (d *DriverHandler) InsertDriverLog(c *fiber.Ctx) error {
	var driver entities.DriverEntity
	if err := c.BodyParser(&driver); err != nil {
		return err
	}

	res, err := d.pc.InsertDriverLog(c.Context(), &types.InsertDriverLogRequest{
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

func (d *DriverHandler) GetDriverLogs(c *fiber.Ctx) error {

	driverId := c.Params("id")

	parsedId, err := strconv.Atoi(driverId)

	if err != nil {
		// Handle the error if the conversion fails
		fmt.Println("Error converting driver ID to integer:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid driver ID format",
		})
	}

	res, err := d.dc.GetDriverLogs(c.Context(), &types.GetDriverLogRequest{
		DriverId: int32(parsedId),
	})

	if err != nil {
		return err
	}

	if !res.DataFound {
		return c.JSON(fiber.Map{})
	}

	var driverLog = entities.DriverLog{
		DriverId: int(res.DriverId),
		Distance: float64(res.DistanceTravelled),
	}

	return c.JSON(driverLog)
}
