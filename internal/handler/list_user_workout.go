package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/response"
	"github.com/gofiber/fiber/v2"
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
