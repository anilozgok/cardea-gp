package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func LogoutHandler(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.ClearCookie("jwt")
		return c.SendString("Logged out successfully!")
	}
}
