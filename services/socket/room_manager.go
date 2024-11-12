// controllers/socket/room_manager.go
package socket

import (
	"encoding/json"
	"go-websocket-fiber/models/types"
	"log"
	"sync"

	"github.com/gofiber/websocket/v2"
)

// RoomManager manages rooms with users and their WebSocket connections
type RoomManager struct {
	rooms map[string]map[string]*websocket.Conn // Map of roomID to userID to WebSocket connection
	mu    sync.Mutex                            // Mutex for synchronizing access to the rooms map
}

// globalRoomManager is a singleton instance of RoomManager
var globalRoomManager = NewRoomManager()

// NewRoomManager initializes and returns a RoomManager instance
func NewRoomManager() *RoomManager {
	return &RoomManager{
		rooms: make(map[string]map[string]*websocket.Conn),
	}
}

// GetRoomManager returns the singleton instance of RoomManager
func GetRoomManager() *RoomManager {
	return globalRoomManager
}

// AddConnection adds a WebSocket connection for a specific user in a specific room
func (rm *RoomManager) AddConnection(roomID, userID string, conn *websocket.Conn) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	// Initialize the room if it doesn't exist
	if _, exists := rm.rooms[roomID]; !exists {
		rm.rooms[roomID] = make(map[string]*websocket.Conn)
	}

	// Add or replace the user’s connection in the room
	if existingConn, exists := rm.rooms[roomID][userID]; exists {
		existingConn.Close() // Close the existing connection if it exists
	}

	rm.rooms[roomID][userID] = conn
}

// RemoveConnection removes a WebSocket connection for a specific user in a specific room
func (rm *RoomManager) RemoveConnection(roomID, userID string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	// Check if the room and user connection exist
	if room, roomExists := rm.rooms[roomID]; roomExists {
		if conn, userExists := room[userID]; userExists {
			conn.Close() // Close the WebSocket connection
			delete(room, userID) // Remove the user’s connection from the room
		}

		// If the room has no more users, delete the room
		if len(room) == 0 {
			delete(rm.rooms, roomID)
		}
	}
}

// BroadcastToRoom sends a message to all users in a specified room
func (rm *RoomManager) BroadcastToRoom(roomID string, message types.RoomMessage) {
	rm.mu.Lock()
	room, exists := rm.rooms[roomID]
	rm.mu.Unlock()

	if !exists {
		log.Printf("No active connections in room: %s\n", roomID)
		return
	}

	// Marshal the message to JSON
	response, err := json.Marshal(message)
	if err != nil {
		log.Println("JSON marshal error:", err)
		return
	}

	// Send the message to each user connection in the room
	log.Println("BroadcastToRoom", roomID)
	for _, conn := range room {
		if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
			log.Println("Write error:", err)
			conn.Close()
			// Optionally remove the user on error:
			// rm.RemoveConnection(roomID, userID) -- if you keep track of userID here
		}
	}
}
