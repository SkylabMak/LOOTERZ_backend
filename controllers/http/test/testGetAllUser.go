package httpTest

import (
	dbClient "LOOTERZ_backend/config/database"
	"LOOTERZ_backend/models/modelsDB"
	"context"

	// "go-websocket-fiber/models"
	"LOOTERZ_backend/prisma/db"
	"log"

	"github.com/gofiber/fiber/v2"
)

func TestGetAllUser(c *fiber.Ctx) error {
	// token := c.Cookies("token")
	// log.Printf("Token cookie: %s\n", token)
	prismaDB := dbClient.GetPrismaDBClient()
	ctx := context.Background()

	// Query users where userID starts with "user"
	users, err := prismaDB.User.FindMany(
		db.User.UserID.Contains("user"),
	).Exec(ctx)

	if err != nil {
		log.Println("Error finding users:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Unable to retrieve users",
			"details": err.Error(),
		})
	}

	return c.JSON(users)
}

func TestGetAllUserGROM(c *fiber.Ctx) error {
	var user []modelsDB.User
	err := dbClient.DB.Where("RoomID LIKE ?", "room%").Find(&user).Error

	// Check if there was an error in querying the database
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to retrieve users",
			"details": err.Error(),
		})
	}

	// Return the users if the query was successful
	return c.JSON(user)
}