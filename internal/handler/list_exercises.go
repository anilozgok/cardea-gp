package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"github.com/anilozgok/cardea-gp/internal/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type ListExercisesHandler struct {
	repo database.Repository
}

func NewListExercisesHandler(repo database.Repository) *ListExercisesHandler {
	return &ListExercisesHandler{
		repo: repo,
	}
}

func (h *ListExercisesHandler) Handle(c *fiber.Ctx) error {
	exercisesList, err := h.repo.ListExercises(c.Context())
	if err != nil {
		zap.L().Error("error while listing exercises", zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	exercises := lo.Map(exercisesList, func(e entity.Exercise, _ int) response.ExerciseResponse {
		return response.ExerciseResponse{
			ExerciseId: e.ID,
			BodyPart:   e.BodyPart,
			Equipment:  e.Equipment,
			Gif:        e.Gif,
			Name:       e.Name,
			Target:     e.Target,
		}
	})

	return c.JSON(response.ExerciseListResponse{Exercises: exercises})
}
