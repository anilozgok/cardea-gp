package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UpdateDietHandler struct {
	repo database.Repository
}

func NewUpdateDietHandler(repo database.Repository) *UpdateDietHandler {
	return &UpdateDietHandler{repo: repo}
}

func (h *UpdateDietHandler) Handle(c *fiber.Ctx) error {
	req := new(request.UpdateDietRequest)
	if err := c.BodyParser(req); err != nil {
		zap.L().Error("error while parsing request body", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	if err := validators.Validate(req); err != nil {
		zap.L().Error("error while validating request body", zap.Error(err))
		c.Status(fiber.StatusBadRequest)
		return err
	}

	diet, err := h.repo.GetDietByID(c.Context(), req.ID)
	if err != nil {
		zap.L().Error("error while getting diet", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	diet.Name = req.Name
	diet.Meals = req.Meals

	if err := h.repo.UpdateDiet(c.Context(), diet); err != nil {
		zap.L().Error("error while updating diet", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
