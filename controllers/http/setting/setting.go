package httpSetting

import (
	"github.com/gofiber/fiber/v2"
)

func CreateRoom(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{})
}


func ChangeSetting(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{})
}