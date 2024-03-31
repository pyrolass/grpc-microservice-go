package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pyrolass/grpc-microservice-go/gateway/entities"
	types "github.com/pyrolass/grpc-microservice-go/proto"
	"google.golang.org/grpc"
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

	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	gClient := types.NewDriverServiceClient(conn)

	_, err = gClient.InsertDriverLog(c.Context(), &types.InsertDriverLogRequest{
		DriverId: int32(driver.DriverId),
		Lat:      driver.Lat,
		Lon:      driver.Lon,
	})

	if err != nil {
		return err
	}

	return c.JSON(map[string]string{
		"message": "destination received",
	})
}
