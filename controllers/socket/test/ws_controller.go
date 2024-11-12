// controllers/socket/ws_controller.go
package socketTest

import (
	"LOOTERZ_backend/models/types"
	"LOOTERZ_backend/services/socket"
	"log"
	"time"

	"github.com/gofiber/websocket/v2"
)

func WebSocketHandler(c *websocket.Conn) {
	roomID := c.Params("roomID")
	token := c.Cookies("token")
	log.Printf("Token cookie: %s\n", token)
	userID := token // Assuming userID is derived from the token
	log.Printf("User %s joined room: %s\n", userID, roomID)

	// Get the singleton instance of RoomManager
	roomManager := socket.GetRoomManager()

	// Add the connection for the user in the room
	roomManager.AddConnection(roomID, userID, c)

	// Channel to signal when the connection is closed
	done := make(chan struct{})

	// Start a goroutine to listen for disconnections
	go func() {
		defer close(done) // Close the done channel when the goroutine ends

		for {
			// Listen for messages to detect disconnection
			if _, _, err := c.ReadMessage(); err != nil {
				log.Println("Read error or client disconnected:", err)
				roomManager.RemoveConnection(roomID, userID)
				break // Exit the loop on error, which will trigger the done channel
			}
		}
	}()

	// Periodic broadcast loop, which stops when done is signaled
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Broadcast message every 5 seconds
			log.Printf("broadcast loop from user %s to room %s\n", userID, roomID)
			roomMessage := generateMockRoomMessage()
			roomManager.BroadcastToRoom(roomID, roomMessage)
		case <-done:
			// Stop the ticker loop when the connection is closed
			log.Printf("Stopping broadcast loop for user %s in room %s\n", userID, roomID)
			return
		}
	}
}

// generateMockRoomMessage generates a mock message for testing purposes
func generateMockRoomMessage() types.RoomMessage {
	return types.RoomMessage{
		Players:      []types.Player{{Name: "Test Player", Ready: true, ImgURL: "http://example.com/img.jpg"}},
		MaxPlayers:   5,
		PlayersCount: 1,
		Time:         120,
		RoomPassword: "secret",
		RoomName:     "Room 1234",
	}
}
