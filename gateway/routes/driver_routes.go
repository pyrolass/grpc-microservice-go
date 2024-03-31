package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pyrolass/grpc-microservice-go/gateway/handlers"
	types "github.com/pyrolass/grpc-microservice-go/proto"
)

func DriverRoutes(router fiber.Router, gClient types.DriverServiceClient) {

	driverHandle := handlers.NewDriverHandler(gClient)
	router.Post("/drivers", driverHandle.InsertDriverLog)

}
