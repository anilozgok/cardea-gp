package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"github.com/anilozgok/cardea-gp/internal/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type ListUsersHandler struct {
	repo database.Repository
}

func NewListUsersHandler(repo database.Repository) *ListUsersHandler {
	return &ListUsersHandler{repo: repo}
}

func (h *ListUsersHandler) Handle(c *fiber.Ctx) error {
	coachId := c.Locals("userId").(uint)

	workouts, err := h.repo.ListWorkoutByCoachId(c.Context(), coachId)
	if err != nil {
		zap.L().Error("error while listing workouts", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	userIds := lo.Map(workouts, func(w entity.Workout, _ int) uint {
		return w.UserId
	})

	users, err := h.repo.GetUsers(c.Context())
	if err != nil {
		zap.L().Error("error while listing users", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	usersFiltered := lo.Map(lo.Filter(users, func(u entity.User, i int) bool {
		return lo.Contains(userIds, u.ID)
	}), func(u entity.User, _ int) response.UserResponse {
		return response.UserResponse{
			UserId:    u.ID,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		}
	})

	return c.JSON(response.UserListResponse{Users: usersFiltered})
}
