package httpTest

import (
	"LOOTERZ_backend/utils"

	"log"

	"github.com/gofiber/fiber/v2"
)

func TestFuntion(c *fiber.Ctx) error {
	token := c.Cookies("token")
	log.Printf("Token cookie: %s\n", token)

	testUser, _ := utils.GenerateJWT("user06")
	log.Printf("encrypt  %s\n", testUser)
	userID, _ := utils.DecodeJWT(testUser)
	log.Printf("decrypt Token cookie: %s\n", userID)
	userUUID, _ := utils.GenerateUUID()
	log.Printf("uuid: %s\n", userUUID)

	utils.SetCookieToken(c, testUser)
	return c.JSON(fiber.Map{})
}