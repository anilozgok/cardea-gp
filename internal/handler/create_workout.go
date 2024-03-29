package handler

import (
	"errors"
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/entities"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type CreateWorkoutHandler struct {
	repo database.Repository
}

func NewCreateWorkoutHandler(repo database.Repository) *CreateWorkoutHandler {
	return &CreateWorkoutHandler{
		repo: repo,
	}
}

func (h *CreateWorkoutHandler) Handle(c *fiber.Ctx) error {
	req := new(request.CreateWorkoutRequest)
	if err := c.BodyParser(req); err != nil {
		zap.L().Error("error while parsing request body", zap.Error(err))
		c.Status(fiber.StatusBadRequest)
		return err
	}

	if err := validators.ValidateCreateWorkoutRequest(req); err != nil {
		zap.L().Error("error while validating request body", zap.Error(err))
		c.Status(fiber.StatusBadRequest)
		return err
	}

	maybeUser, err := h.repo.GetUserById(c.Context(), req.UserId)
	if err != nil {
		zap.L().Error("error while checking user existence", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	if maybeUser == nil {
		zap.L().Error("user not found", zap.Error(err))
		c.Status(fiber.StatusBadRequest)
		return errors.New("user not exists")
	}

	workout := entities.Workout{
		UserId:      req.UserId,
		Name:        req.Name,
		Description: req.Description,
		Area:        req.Area,
		Rep:         req.Rep,
		Sets:        req.Sets,
	}

	if err = h.repo.CreateWorkout(c.Context(), &workout); err != nil {
		zap.L().Error("error while creating workout", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return nil
}
