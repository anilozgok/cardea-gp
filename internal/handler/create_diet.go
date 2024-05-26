package handler

import (
	"errors"
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CreateDietHandler struct {
	repo database.Repository
}

func NewCreateDietHandler(repo database.Repository) *CreateDietHandler {
	return &CreateDietHandler{
		repo: repo,
	}
}

func (h *CreateDietHandler) Handle(c *fiber.Ctx) error {
	req := new(request.CreateDietRequest)
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

	coachId := c.Locals("userId").(uint)

	maybeUser, err := h.repo.GetUserById(c.Context(), req.UserId)
	if err != nil {
		zap.L().Error("error while getting user", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	if maybeUser == nil {
		zap.L().Error("user not found")
		c.Status(fiber.StatusNotFound)
		return errors.New("user not found")
	}

	diet := entity.Diet{
		UserId:      req.UserId,
		CoachId:     coachId,
		Meal:        req.Meal,
		Description: req.Description,
		Calories:    req.Calories,
		Protein:     req.Protein,
		Carbs:       req.Carbs,
		Fat:         req.Fat,
	}

	if err = h.repo.CreateDiet(c.Context(), &diet); err != nil {
		zap.L().Error("error while creating diet", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
