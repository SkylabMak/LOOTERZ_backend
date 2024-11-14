package modelsDB

import "time"

type Room struct {
	RoomID          string    `gorm:"primaryKey;column:roomID" `
	RoomStatus      int       `gorm:"column:roomStatus"`
	RoomName        string    `gorm:"column:roomName"`
	MaxPlayerAmount int       `gorm:"column:maxPlayerAmount"`
	TimePerTurn     int       `gorm:"column:timePerTurn"`
	PrivateStatus   bool      `gorm:"column:privateStatus"`
	Password        string    `gorm:"column:password"`
	CreatedAt       time.Time `gorm:"column:created_at"`

	Users []User `gorm:"foreignKey:RoomID" json:"users"`
}

func (Room) TableName() string {
	return "Room"
}