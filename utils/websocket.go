package utils

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all connections (for testing purposes)
	},
}

func HandleWebSocketConnections(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return nil, err
	}
	return conn, nil
}
