package types

type ErrorCode string

type Player struct {
	Name   string `json:"name"`
	Ready  bool   `json:"ready"`
	ImgURL string `json:"imgURL"`
}

// RoomMessage represents the structure of a WebSocket message for room details
type RoomMessage struct {
	Players      []Player `json:"players"`
	MaxPlayers   int      `json:"maxPlayers"`
	PlayersCount int      `json:"Players"`
	Time         int      `json:"time"`
	RoomPassword string   `json:"roomPassword"`
	RoomName     string   `json:"roomName"`
}

type RoomResponse struct {
	RoomName       string `json:"roomName"`
	RoomID         string `json:"roomID"`
	NumberPlayers  int    `json:"NumberPlayers"`
	CurrentPlayers int    `json:"currentPlayes"`
	Time           int    `json:"time"`
	PrivateStatus  bool   `json:"privateStatus"`
}