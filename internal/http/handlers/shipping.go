package handlers

import (
	usecases "github.com/bruandreo/shipping-service/useCases"
	"github.com/gofiber/fiber/v2"
)

func GetShippings(c *fiber.Ctx) error {
	shippings := usecases.GetShippings()

	if len(shippings) <= 0 {
		return c.Status(fiber.StatusNotFound).JSON(map[string]string{
			"message": "Not found Shippings",
		})
	}

	return c.Status(fiber.StatusOK).JSON(shippings)
}
