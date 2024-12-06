package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx) (string, error) {
	cookie := c.Cookies("token")
	if cookie == "" {
		return "", errors.New("Unautorized")

	}
	userId, err := DecodeJWT(cookie)
	if err != nil {
		return "", err
	}

	return userId, nil

}
