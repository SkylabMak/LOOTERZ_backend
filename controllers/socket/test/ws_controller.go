// ws_controller.go
package socketTest

import (
	"encoding/json"
	"go-websocket-fiber/models/types"
	"log"
	"math/rand"
	"time"

	"github.com/gofiber/websocket/v2"
)

// Handler manages WebSocket connections
func WebSocketHandler(c *websocket.Conn) {
	log.Println("socket run")
	// Close the WebSocket connection when the function ends
	defer c.Close()

	// Create a ticker that ticks every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	// Use a for range loop to iterate over the ticker.C channel
	for range ticker.C {
		// Generate mock RoomMessage data
		roomMessage := generateMockRoomMessage()

		// Encode RoomMessage data to JSON
		response, err := json.Marshal(roomMessage)
		if err != nil {
			log.Println("JSON marshal error:", err)
			continue
		}

		// Send JSON data to client
		if err = c.WriteMessage(websocket.TextMessage, response); err != nil {
			log.Println("Write error:", err)
			return
		}
		log.Println("socket loop run")
	}

}


// generateMockRoomMessage generates mock RoomMessage data
func generateMockRoomMessage() types.RoomMessage {
	players := []types.Player{
		{Name: "Alice", Ready: rand.Intn(2) == 1, ImgURL: "http://example.com/alice.jpg"},
		{Name: "Bob", Ready: rand.Intn(2) == 1, ImgURL: "http://example.com/bob.jpg"},
		{Name: "Charlie", Ready: rand.Intn(2) == 1, ImgURL: "http://example.com/charlie.jpg"},
	}

	return types.RoomMessage{
		Players:      players,
		MaxPlayers:   5,
		PlayersCount: len(players),
		Time:         rand.Intn(120), // Random time value
		RoomPassword: "mockPassword",
		RoomName:     "Mock Room",
	}
}