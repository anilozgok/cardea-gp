package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type UpdateWorkoutHandler struct {
	repo database.Repository
}

func NewUpdateWorkoutHandler(repo database.Repository) *UpdateWorkoutHandler {
	return &UpdateWorkoutHandler{
		repo: repo,
	}
}

func (h *UpdateWorkoutHandler) Handle(c *fiber.Ctx) error {
	req := new(request.UpdateWorkout)
	if err := c.BodyParser(req); err != nil {
		zap.L().Error("error while parsing request body", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err := validators.Validate(req); err != nil {
		zap.L().Error("error while validating request body", zap.Error(err))
		return c.SendStatus(fiber.StatusBadRequest)
	}

	workout, err := h.repo.GetWorkoutById(c.Context(), req.WorkoutId)
	if err != nil {
		zap.L().Error("error while getting workout", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if workout == nil {
		zap.L().Error("workout not found", zap.Uint("workoutId", req.WorkoutId))
		return c.SendStatus(fiber.StatusNotFound)
	}

	workout.Rep = req.Rep
	workout.Sets = req.Sets

	if err = h.repo.UpdateWorkout(c.Context(), *workout); err != nil {
		zap.L().Error("error while updating workout", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
