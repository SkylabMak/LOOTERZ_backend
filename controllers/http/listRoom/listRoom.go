package httpListRoom

import (
	gormDB "LOOTERZ_backend/config/database"
	"LOOTERZ_backend/models/types"
	"LOOTERZ_backend/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetListRoom(c *fiber.Ctx) error {
	page, errPage := strconv.Atoi(c.Params("page"))
	// Define a slice to hold the rooms with user count
	if errPage != nil || page < 1 {
		page = 1
	}
	const limit = 20
	offset := (page - 1) * limit
	var roomsWithCount []struct {
		RoomID          string `gorm:"primaryKey;column:roomID" `
		RoomName        string `gorm:"column:roomName"`
		MaxPlayerAmount int    `gorm:"column:maxPlayerAmount"`
		TimePerTurn     int    `gorm:"column:timePerTurn"`
		PrivateStatus   bool   `gorm:"column:privateStatus"`
		CurrentPlayers  int
	}

	// Query to join Room and User tables, counting user per room
	err := gormDB.DB.Table("room").
		Select("room.roomID, room.roomName, room.maxPlayerAmount, room.timePerTurn, room.privateStatus, COUNT(user.userID) AS current_players").
		Joins("LEFT JOIN user ON user.roomID = room.roomID").
		Where("room.roomID LIKE ?", "room%").
		Group("room.roomID").
		Order("room.created_at DESC"). // Order by creation date
		Limit(limit).
		Offset(offset).
		Find(&roomsWithCount).Error

	// log.Println(roomsWithCount)
	if err != nil {
		return utils.FullErrorResponse(c, 500, utils.ErrInternal, "Unable to retrieve rooms", err)
	}

	// Map results to RoomResponse struct
	var responseRooms []types.RoomResponse
	for _, room := range roomsWithCount {
		responseRooms = append(responseRooms, types.RoomResponse{
			RoomID:         room.RoomID,
			RoomName:       room.RoomName,
			NumberPlayers:  room.MaxPlayerAmount,
			CurrentPlayers: room.CurrentPlayers,
			Time:           room.TimePerTurn,
			PrivateStatus:  room.PrivateStatus,
		})
	}

	if(responseRooms == nil){
		return c.JSON([]interface{}{})
	}

	return c.JSON(responseRooms)
}

func EnterRoom(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{})
}
