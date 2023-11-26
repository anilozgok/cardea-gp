package handlers

import "github.com/gofiber/fiber/v2"

func NewGetUsersHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		//TODO:: GetUsers logic
		list := []string{"user1", "user2", "user3"}

		return c.JSON(list)
	}
}
