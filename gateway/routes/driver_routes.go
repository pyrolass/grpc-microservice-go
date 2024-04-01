package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pyrolass/grpc-microservice-go/gateway/handlers"
	types "github.com/pyrolass/grpc-microservice-go/proto"
)

func DriverRoutes(router fiber.Router, pc types.DriverMessageServiceClient, dc types.DriverQueryServiceClient) {

	driverHandle := handlers.NewDriverHandler(pc, dc)

	router.Post("/drivers", driverHandle.InsertDriverLog)
	router.Get("/drivers/:id", driverHandle.GetDriverLogs)

}
