package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ListFoodsHandler struct {
	repo database.Repository
}

func NewListFoodsHandler(repo database.Repository) *ListFoodsHandler {
	return &ListFoodsHandler{
		repo: repo,
	}
}

func (h *ListFoodsHandler) Handle(c *fiber.Ctx) error {
	foods, err := h.repo.ListFoods(c.Context())
	if err != nil {
		zap.L().Error("error while listing foods", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(foods)
}
