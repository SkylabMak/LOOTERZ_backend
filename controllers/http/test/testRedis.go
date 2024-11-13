package httpTest

import (
	"LOOTERZ_backend/services/socket"
	"context"

	"github.com/gofiber/fiber/v2"
)

// POST handler to send messages
func SendMessageHandler(c *fiber.Ctx) error {
	// Parse JSON body
	var request struct {
		Channel string `json:"channel"`
		Message string `json:"message"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	ctx := context.Background()
	// Publish message to Redis channel
	err := socket.RedisClient.Publish(ctx, request.Channel, request.Message).Err()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to publish message",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "Message published successfully",
		"channel": request.Channel,
		"message": request.Message,
	})
}