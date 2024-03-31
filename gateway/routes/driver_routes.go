package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pyrolass/grpc-microservice-go/gateway/handlers"
)

func DriverRoutes(router fiber.Router) {

	driverHandle := handlers.NewDriverHandler()
	router.Post("/drivers", driverHandle.InsertDriverLog)

}
