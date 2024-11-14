package utils

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse sends a standardized error response
func ErrorResponse(c *fiber.Ctx,statusCode int, message string,  detail string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error":   message,
		"details": detail,
	})
}

func FullErrorResponse(c *fiber.Ctx, statusCode int, message string, err error) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error":   message,
		"details": err.Error(),
	})
}