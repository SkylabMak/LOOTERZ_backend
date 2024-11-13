package httpListRoom

import (
	"github.com/gofiber/fiber/v2"
)

func GetListRoom(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{})
}


func EnterRoom(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{})
}