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

	usersOfCoach, err := h.repo.GetStudentsOfCoach(c.Context(), coachId)
	if err != nil {
		zap.L().Error("error while getting students", zap.Uint("coachId", coachId), zap.Error(err))
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	users := lo.Map(usersOfCoach, func(u entity.User, _ int) response.UserResponse {
		return response.UserResponse{
			UserId:    u.ID,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		}
	})

	return c.JSON(response.UserListResponse{Users: users})
}
