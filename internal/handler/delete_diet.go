package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
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
	dietId := c.Query("dietId")
	if dietId == "" {
		zap.L().Error("dietId not found on the query param")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	id, err := strconv.Atoi(dietId)
	if err != nil {
		zap.L().Error("error while parsing dietId", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err = h.repo.DeleteDiet(c.Context(), uint(id)); err != nil {
		zap.L().Error("error while deleting diet", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
