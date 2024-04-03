package main

import (
	"github.com/bruandreo/shipping-service/internal/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		AppName:       "Shipping-service",
	})

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${method} ${path} ${status} ${latency}\n",
	}))

	http.InitializeRoutes(app)

	app.Listen(":3000")
}
