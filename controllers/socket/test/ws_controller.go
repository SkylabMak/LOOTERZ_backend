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
	//chan struct{} is a channel of empty structs (struct{}), and it’s often used as a signal-only channel because an empty struct in Go (struct{}) is a zero-sized type
	done := make(chan struct{})

	// Start a goroutine to listen for disconnections
	//The go keyword before func() tells Go to execute the function in a new, concurrent goroutine. This way, the code inside this function runs concurrently, independent of the main function.
	go func() {
		/*
			1. Start `for` loop
			2. Wait for message in `c.ReadMessage()` (blocks until a message arrives)
			3. Message arrives or error occurs
			4. Check error:
			   - If no error, go back to step 2 and wait for the next message.
			   - If error, handle disconnection:
			       a. Log the disconnection and remove the connection from `RoomManager`.
			       b. Break out of the loop and trigger `defer close(done)` to signal disconnection.
			5. Goroutine ends after signaling disconnection.
		*/
		// Close the done channel when the goroutine ends. to run only when the function ends,
		defer close(done)

		for {
			// Listen for messages to detect disconnection
			// that waits for a message from the WebSocket connection.
			if _, _, err := c.ReadMessage(); err != nil {
				//will return an error, allowing us to detect that disconnectio
				//It doesn’t skip, pass, or continue; it simply waits for c.ReadMessage() to finish.
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
		/*
		   Execution Flow:
		   1. Start `for` loop and enter `select` statement.
		   2. `select` "waits" for either `ticker.C` or `done` to be ready.
		   3. If `ticker.C` sends a tick (every 5 seconds):
		   4. If `done` is closed (indicating a disconnection):
		*/
		//ticker and done is channel
		//select is "wait" for any case then execute then back to "wait"
		select {
		case <-ticker.C: //This value goes into the channel and becomes "available," allowing the case <-ticker.C: to execute in every 5 second
			// Broadcast message every 5 seconds
			log.Printf("broadcast loop from user %s to room %s\n", userID, roomID)
			roomMessage := generateMockRoomMessage()
			roomManager.BroadcastToRoom(roomID, roomMessage)
		case <-done: //When done is closed, it becomes "ready" in the select statement
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
