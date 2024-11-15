// utils/content_type.go
package security

import "github.com/gofiber/fiber/v2"

// CheckContentType checks if the Content-Type of the request matches the expected type.
// If it doesn't match, it returns an error response.
// CheckContentType returns true if the Content-Type matches the expected type, otherwise false.
func CheckContentType(c *fiber.Ctx, expectedType string) bool {
	return c.Get("Content-Type") == expectedType
}

// CheckContentJSONType returns true if the Content-Type is "application/json", otherwise false.
func CheckContentJSONType(c *fiber.Ctx) bool {
	return c.Get("Content-Type") == "application/json"
}