package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strconv"
)

type DeleteWorkoutHandler struct {
	repo database.Repository
}

func NewDeleteWorkoutHandler(repo database.Repository) *DeleteWorkoutHandler {
	return &DeleteWorkoutHandler{
		repo: repo,
	}
}

func (h *DeleteWorkoutHandler) Handle(c *fiber.Ctx) error {
	workoutId := c.Query("workoutId")
	if workoutId == "" {
		zap.L().Error("workoutId not found on the query param")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	id, err := strconv.Atoi(workoutId)
	if err != nil {
		zap.L().Error("error while parsing workoutId", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err = h.repo.DeleteWorkout(c.Context(), uint(id)); err != nil {
		zap.L().Error("error while deleting workout", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
