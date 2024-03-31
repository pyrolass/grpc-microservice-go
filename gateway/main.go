package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pyrolass/grpc-microservice-go/gateway/common"
	"github.com/pyrolass/grpc-microservice-go/gateway/middleware"
	"github.com/pyrolass/grpc-microservice-go/gateway/routes"
	types "github.com/pyrolass/grpc-microservice-go/proto"
	"google.golang.org/grpc"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if apiError, ok := err.(common.ApiError); ok {
			return c.Status(apiError.Code).JSON(
				map[string]any{
					"msg": apiError.Message,
				},
			)
		}

		return c.Status(500).JSON(
			map[string]any{
				"msg": err.Error(),
			},
		)
	},
}

func main() {
	app := fiber.New(config)

	router := app.Group("api/v1", middleware.LogRequest)

	router.Get(
		"/ping",
		func(c *fiber.Ctx) error {
			return c.JSON(
				map[string]string{
					"message": "pong",
				},
			)
		},
	)

	// driver routes grpc
	conn, err := grpc.Dial("localhost:3001", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	gClient := types.NewDriverMessageServiceClient(conn)

	routes.DriverRoutes(router, gClient)

	log.Fatal(app.Listen(":3000"))
}
