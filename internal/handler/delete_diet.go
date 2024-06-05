package handler

import (
	"strconv"

	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type DeleteDietHandler struct {
	repo database.Repository
}

func NewDeleteDietHandler(repo database.Repository) *DeleteDietHandler {
	return &DeleteDietHandler{
		repo: repo,
	}
}

func (h *DeleteDietHandler) Handle(c *fiber.Ctx) error {
	dietIdStr := c.Query("diet_id")
	if dietIdStr == "" {
		return c.Status(fiber.StatusBadRequest).SendString("diet_id is required")
	}

	dietId, err := strconv.ParseUint(dietIdStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid diet_id format")
	}

	if err := h.repo.DeleteDiet(c.Context(), uint(dietId)); err != nil {
		zap.L().Error("error while deleting diet", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("error while deleting diet")
	}

	return c.SendStatus(fiber.StatusOK)
}
