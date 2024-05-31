package handler

import (
	"strconv"

	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ListDietHandler struct {
	repo database.Repository
}

func NewListDietHandler(repo database.Repository) *ListDietHandler {
	return &ListDietHandler{
		repo: repo,
	}
}

func (h *ListDietHandler) Handle(c *fiber.Ctx) error {
	userIdStr := c.Query("user_id")
	if userIdStr == "" {
		return c.Status(fiber.StatusBadRequest).SendString("user_id is required")
	}

	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid user_id format")
	}

	diets, err := h.repo.ListDiets(c.Context(), uint(userId))
	if err != nil {
		zap.L().Error("error while listing diets", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("error while listing diets")
	}

	return c.JSON(diets)
}
