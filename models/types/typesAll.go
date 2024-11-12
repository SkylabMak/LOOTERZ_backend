package types

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