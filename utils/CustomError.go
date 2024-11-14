package utils

import (
	"github.com/gofiber/fiber/v2"
)

// ErrorResponse sends a standardized error response
func ErrorResponse(c *fiber.Ctx, statusCode int, errorCode string, message string, detail string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"code":    errorCode,
		"error":   message,
		"details": detail,
	})
}

func FullErrorResponse(c *fiber.Ctx, statusCode int, errorCode string, message string, err error) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"code":    errorCode,
		"error":   message,
		"details": err.Error(),
	})
}

func CustomErrorResponse(c *fiber.Ctx, statusCode int, errorCode string, header string, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"code": errorCode,
		header: message,
	})
}
