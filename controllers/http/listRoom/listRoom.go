package httpListRoom

import (
	gormDB "LOOTERZ_backend/config/database"
	"LOOTERZ_backend/models/modelsDB"
	"LOOTERZ_backend/models/types"
	"LOOTERZ_backend/utils"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetListRoom(c *fiber.Ctx) error {
	page, errPage := strconv.Atoi(c.Params("page"))

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

	
	err := gormDB.DB.Table("room").
		Select("room.roomID, room.roomName, room.maxPlayerAmount, room.timePerTurn, room.privateStatus, COUNT(user.userID) AS current_players").
		Joins("LEFT JOIN user ON user.roomID = room.roomID").
		Where("room.roomID LIKE ?", "room%").
		Group("room.roomID").
		Order("room.created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&roomsWithCount).Error

	
	if err != nil {
		return utils.FullErrorResponse(c, 500, utils.ErrInternal, "Unable to retrieve rooms", err)
	}

	
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
	var request struct {
		RoomID string `json:"roomID"`
		Password string `json:"password"`
	}

	token := c.Cookies("token")
	userID, errorToken := utils.DecodeJWT(token)
	// Check if there was an error decoding the JWT token
	log.Print(userID)
	if errorToken != nil {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, utils.ErrBadReq, "Invalid or expired token", "Invalid or expired token")
	}

	if err := c.BodyParser(&request); err != nil {
		return utils.ErrorResponse(c,400,utils.ErrBadReq   ,"Bad request","request body miss match")
	}

	var room modelsDB.Room
	if err := gormDB.DB.First(&room, "roomID = ?", request.RoomID).Error; err != nil {
		// Return an error if the room is not found
		return utils.ErrorResponse(c, fiber.StatusNotFound, utils.ErrNotFound, "Room not found", "Room not found")
	}
	return c.JSON(room)
}
