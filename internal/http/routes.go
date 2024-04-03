package http

import (
	"github.com/bruandreo/shipping-service/internal/http/handlers"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	app.Get("/shippings", handlers.GetShippings)
}
