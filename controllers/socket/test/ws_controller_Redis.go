// ws_controller.go
package socketTest

import (
	"LOOTERZ_backend/services/socket"
	"context"
	"log"

	"github.com/gofiber/websocket/v2"
)

// WebSocketHandler handles WebSocket connections and subscribes to a Redis channel for real-time broadcasting
func WebSocketHandler_redis(c *websocket.Conn) {
	roomID := c.Params("roomID")
	log.Printf("User connected to room: %s\n", roomID)

	// Subscribe to the Redis channel for the room
	subscriber := socket.RedisClient.Subscribe(context.Background(), roomID)
	channel := subscriber.Channel()

	// Goroutine to listen to Redis Pub/Sub messages and send them to the WebSocket client
	go func() {
		for msg := range channel {
			if err := c.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
				log.Printf("Error sending message to WebSocket: %v", err)
				return
			}
		}
	}()

	// Ensure subscription is cleaned up when connection is closed
	defer func() {
		subscriber.Close()
		log.Printf("User disconnected from room: %s\n", roomID)
	}()

	// Listen for incoming messages from WebSocket client and publish them to the Redis channel
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("Read error or client disconnected:", err)
			break
		}

		// Publish the message to the Redis channel for broadcasting to all clients in the room
		if err = socket.PublishToRoom(roomID, string(msg)); err != nil {
			log.Printf("Error publishing to room %s: %v", roomID, err)
		}
	}
}
