package httpHome

import (
	"github.com/gofiber/fiber/v2"
)

func Enter(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{})
}