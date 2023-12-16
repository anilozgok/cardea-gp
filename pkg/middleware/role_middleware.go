package middleware

import (
	"github.com/anilozgok/cardea-gp/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func RoleUser(ctx *fiber.Ctx) error {
	role := ctx.Locals("role").(string)

	if role != utils.ROLE_USER {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}

func RoleAdmin(ctx *fiber.Ctx) error {
	role := ctx.Locals("role").(string)

	if role != utils.ROLE_ADMIN {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}

func RoleCoach(ctx *fiber.Ctx) error {
	role := ctx.Locals("role").(string)

	if role != utils.ROLE_COACH {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	return ctx.Next()
}
