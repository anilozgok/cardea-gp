package handlers

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/gofiber/fiber/v2"
	"time"
)

type LogoutHandler struct {
	repo database.Repository
}

func NewLogoutHandler(repo database.Repository) *LogoutHandler {
	return &LogoutHandler{
		repo: repo,
	}
}

func (h *LogoutHandler) Handle(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:    "authSession",
		Value:   "",
		Expires: time.Now().Add(-time.Hour * 24),
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}
