package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/response"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ListCoachWorkoutHandler struct {
	repo database.Repository
}

func NewListCoachWorkoutHandler(repo database.Repository) *ListCoachWorkoutHandler {
	return &ListCoachWorkoutHandler{
		repo: repo,
	}
}

func (h *ListCoachWorkoutHandler) Handle(c *fiber.Ctx) error {
	coachId := c.Locals("userId").(uint)

	workouts, err := h.repo.ListWorkoutByCoachId(c.Context(), coachId)
	if err != nil {
		zap.L().Error("error while listing workouts", zap.Error(err), zap.Uint("coachId", coachId))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	res := make([]response.WorkoutResponse, 0)
	for _, w := range workouts {
		res = append(res, response.WorkoutResponse{
			UserId:      w.UserId,
			CoachId:     w.CoachId,
			Name:        w.Name,
			Description: w.Description,
			Area:        w.Area,
			Rep:         w.Rep,
			Sets:        w.Sets,
		})
	}

	return c.JSON(response.WorkoutListResponse{Workouts: res})
}
