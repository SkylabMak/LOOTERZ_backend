package httpHome

import (
	"LOOTERZ_backend/models/modelsDB"
	"LOOTERZ_backend/utils"

	"github.com/gofiber/fiber/v2"
)

func Enter(c *fiber.Ctx) error {

	// ต้องใส่ & ใน parameter BodyParser ถ้าใช้แบบนี้
	// var user struct {
	// 	Name string `json:"name"`
	// }

	type UserReq struct {
		Name string `json:"name"`
	}
	// user := new(UserReq) // เหมือนกันกับ &UserReq{}
	user := &UserReq{}
	if err := c.BodyParser(user); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, utils.ErrContentType, "Err Content Type", "Content Type must be application/json")
	}

	// Ensure user name is provided
	if user.Name == "" {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "400", "Invalid Input", "User name is required")
	}

	// Generate UUID for new user
	uuid, err := utils.GenerateUUID()
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, utils.ErrInternal, "Generate UUID Error", "Failed to generate user UUID")
	}

	newUser := &modelsDB.User{
		UserID:   uuid,
		UserName: user.Name,
	}

	// Insert user into the database
	if err := createUser(newUser); err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, utils.ErrDatabaseConnection, "Database Error", "Failed to create user")
	}

	// Generate JWT token for new user
	jwtToken, err := utils.GenerateJWT(newUser.UserID)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, utils.ErrInternal, "Jwt Token Error", "Failed to generate JWT token")
	}

	// Set JWT token as cookie
	utils.SetCookieToken(c, jwtToken)

	// Return success response
	return c.Send(nil)
}

func EnterCheckStatus(c *fiber.Ctx) error {

	cookie := c.Cookies("token")
	if cookie == "" {
		return utils.ErrorResponse(c, fiber.StatusUnauthorized, "401", "Unautorized", "Failed to Authorization")

	}
	userId, err := utils.DecodeJWT(cookie)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, utils.ErrInternal, "Decode JWT Error", "Failed to decoding jwt")
	}

	status, err := checkUserInRoom(userId)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, utils.ErrInternal, "Check User Status Error", "Failed to check user status")
	}

	return c.JSON(fiber.Map{
		"status": status,
	})
}
