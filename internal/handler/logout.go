package handler

import (
	"github.com/gofiber/fiber/v2"
	"time"
)

type LogoutHandler struct{}

func NewLogoutHandler() *LogoutHandler {
	return &LogoutHandler{}
}

func (h *LogoutHandler) Handle(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "authSession",
		Expires: time.Now(),
	})

	return c.SendStatus(fiber.StatusNoContent)
}
