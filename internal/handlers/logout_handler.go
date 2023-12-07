package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

//TODO:: should we clear the cookie when we logout or we just set cookie value to the empty string and expire date to current time? lets discuss
func LogoutHandler(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.ClearCookie("jwt")
		return c.SendString("Logged out successfully!")
	}
}
