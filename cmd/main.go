package main

import (
	"github.com/bruandreo/shipping-service/internal/http"
	"github.com/bruandreo/shipping-service/pkg/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	config := config.Initialize()

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       config.ServiceName,
	})

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${method} ${path} ${status} ${latency}\n",
	}))

	http.InitializeRoutes(app)

	app.Listen(":" + config.AppPort)
}
