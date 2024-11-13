package httpLobby

import (
	"github.com/gofiber/fiber/v2"
)

func PlayerReady(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{})
}