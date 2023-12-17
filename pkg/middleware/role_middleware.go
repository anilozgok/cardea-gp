package middleware

import (
	"github.com/anilozgok/cardea-gp/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func RoleUser(c *fiber.Ctx) error {
	role := c.Locals("role").(string)

	if role != utils.ROLE_USER {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}

func RoleAdmin(c *fiber.Ctx) error {
	role := c.Locals("role").(string)

	if role != utils.ROLE_ADMIN {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}

func RoleCoach(c *fiber.Ctx) error {
	role := c.Locals("role").(string)

	if role != utils.ROLE_COACH {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
