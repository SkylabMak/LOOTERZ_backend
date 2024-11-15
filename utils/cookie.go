// utils/cookie.go
package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// SetCookie sets a cookie with the specified name, value, and expiration time.
func SetCookie(c *fiber.Ctx, name, value string, expires time.Duration) {
	c.Cookie(&fiber.Cookie{
		Name:     name,
		Value:    value,
		Expires:  time.Now().Add(expires),
		SameSite: "Strict",
		HTTPOnly: true,
		Secure:   true, 
		Path: "/",
	})
}

// SetCookieWithDefaults sets a cookie with default values for name and expiration.
func SetCookieToken(c *fiber.Ctx, value string) {
	// Default parameters
	defaultName := "token"
	defaultExpires := 24 * time.Hour

	// Call SetCookie with default name and expiration
	SetCookie(c, defaultName, value, defaultExpires)
}