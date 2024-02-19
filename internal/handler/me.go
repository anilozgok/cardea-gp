package handler

import (
	"github.com/anilozgok/cardea-gp/internal/model/response"
	"github.com/gofiber/fiber/v2"
)

type MeHandler struct {
}

func NewMeHandler() *MeHandler {
	return &MeHandler{}
}

func (h *MeHandler) Handle(c *fiber.Ctx) error {
	return c.JSON(
		response.MeResponse{
			UserId: c.Locals("userId").(uint32),
			Email:  c.Locals("email").(string),
			Role:   c.Locals("role").(string),
		},
	)
}
