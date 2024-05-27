package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"github.com/anilozgok/cardea-gp/internal/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type ListUserWorkoutHandler struct {
	repo database.Repository
}

func NewListUserWorkoutHandler(repo database.Repository) *ListUserWorkoutHandler {
	return &ListUserWorkoutHandler{
		repo: repo,
	}
}

func (h *ListUserWorkoutHandler) Handle(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	workouts, err := h.repo.ListWorkoutByUserId(c.Context(), userId)
	if err != nil {
		zap.L().Error("error while listing workouts", zap.Error(err), zap.Uint("userId", userId))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	workoutRes := lo.Map(workouts, func(w entity.Workout, _ int) response.WorkoutResponse {
		exercise, err := h.repo.GetExerciseById(c.Context(), w.Exercise)
		if err != nil {
			zap.L().Error("error while listing workouts", zap.Error(err), zap.Uint("userId", userId))
			c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
			return response.WorkoutResponse{}
		}

		return response.WorkoutResponse{
			WorkoutId:    w.ID,
			UserId:       w.UserId,
			CoachId:      w.CoachId,
			Name:         w.Name,
			Description:  w.Description,
			Area:         w.Area,
			Rep:          w.Rep,
			Sets:         w.Sets,
			Exercise:     w.Exercise,
			ExerciseName: exercise.Name,
			Gif:          exercise.Gif,
			Equipment:    exercise.Equipment,
		}
	})

	return c.JSON(response.WorkoutListResponse{Workouts: workoutRes})
}
