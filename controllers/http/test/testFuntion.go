package httpTest

import (
	"LOOTERZ_backend/utils"
	"LOOTERZ_backend/utils/security"

	"log"

	"github.com/gofiber/fiber/v2"
)

func TestFuntion(c *fiber.Ctx) error {
	token := c.Cookies("token")
	log.Printf("Token cookie: %s\n", token)

	testUser, _ := utils.GenerateJWT("user08")
	log.Printf("encrypt  %s\n", testUser)
	userID, _ := utils.DecodeJWT(testUser)
	log.Printf("decrypt Token cookie: %s\n", userID)
	userUUID, _ := utils.GenerateUUID()
	log.Printf("uuid: %s\n", userUUID)

	utils.SetCookieToken(c, testUser)
	return c.JSON(fiber.Map{
		"token": testUser,
	})
}

func TestFuntion02(c *fiber.Ctx) error {
	if !security.CheckContentJSONType(c) {
		return utils.ErrorResponse(c, 400, utils.ErrContentType, "Err Content Type", "Content Type must be application/json")
	}
	var request struct {
		Token   string `json:"token"`
	}

	if err := c.BodyParser(&request); err != nil {
		return utils.ErrorResponse(c, 400, utils.ErrBadReq, "Bad request", "request body miss match")
	}

	token := request.Token
	log.Printf("Token cookie: %s\n", token)

	testUser, _ := utils.GenerateJWT("user08")
	log.Printf("encrypt  %s\n", testUser)
	userID, _ := utils.DecodeJWT(testUser)
	log.Printf("decrypt Token cookie: %s\n", userID)
	userUUID, _ := utils.GenerateUUID()
	log.Printf("uuid: %s\n", userUUID)

	utils.SetCookieToken(c, testUser)
	return c.JSON(fiber.Map{
		"token": testUser,
	})
}