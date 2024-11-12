// ws_controller.go
package socket

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

// WebSocketHandler handles WebSocket connections
func WebSocketHandler(c *websocket.Conn) {
	var (
		messageType int
		msg         []byte
		err         error
	)

	for {
		// Read message from client
		if messageType, msg, err = c.ReadMessage(); err != nil {
			log.Println("Read error:", err)
			break
		}

		// Print received message
		log.Printf("Received message: %s\n", msg)

		// Write message back to client
		if err = c.WriteMessage(messageType, msg); err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
