// controllers/socket/room_manager.go
package socket

import (
	"LOOTERZ_backend/models/types"
	"sync"

	"github.com/gofiber/websocket/v2"
)

// RoomManager manages rooms with users and their WebSocket connections
type RoomManager_socketIO struct {
	rooms map[string]*websocket.Conn // Map of roomID to userID to WebSocket connection
	mu    sync.Mutex                            // Mutex for synchronizing access to the rooms map
}

// globalRoomManager is a singleton instance of RoomManager
// package-level variable that is initialized only once and shared across the program.
var globalRoomManager_socketIO = NewRoomManager_socketIO()

/*
var ExportedVar = "I can be accessed from other packages"
var unexportedVar = "I am only accessible within the example package"
*/
// NewRoomManager initializes and returns a RoomManager instance
func NewRoomManager_socketIO() *RoomManager_socketIO {
	return &RoomManager_socketIO{
		rooms: make(map[string]*websocket.Conn),
		//The types slice, map, and channel require make because they are reference types that involve underlying data structures that need setup:
	}
}

// GetRoomManager returns the singleton instance of RoomManager
func GetRoomManager_socketIO() *RoomManager_socketIO {
	return globalRoomManager_socketIO
}

// AddConnection adds a WebSocket connection for a specific user in a specific room
func (rm *RoomManager_socketIO) AddConnection_socketIO(roomID, userID string, conn *websocket.Conn) {
	
}

// RemoveConnection removes a WebSocket connection for a specific user in a specific room
func (rm *RoomManager_socketIO) RemoveConnection_socketIO(roomID, userID string) {
	
}

// BroadcastToRoom sends a message to all users in a specified room
func (rm *RoomManager_socketIO) BroadcastToRoom_socketIO(roomID string, message types.RoomMessage) {
	
}

// CheckRoomExists checks if a room with the given roomID exists
func (rm *RoomManager_socketIO) CheckRoomExists(roomID string) bool {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	_, exists := rm.rooms[roomID]
	return exists
}

// CheckUserInRoom checks if a user with the given userID is in the specified room
func (rm *RoomManager_socketIO) CheckUserInRoom(roomID, userID string) bool {

	return false
}
