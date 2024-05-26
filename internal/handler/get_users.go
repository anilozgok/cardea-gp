package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"github.com/anilozgok/cardea-gp/internal/model/response"
	"github.com/anilozgok/cardea-gp/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type GetUsersHandler struct {
	repo database.Repository
}

func NewGetUsersHandler(repo database.Repository) *GetUsersHandler {
	return &GetUsersHandler{repo: repo}
}

func (h *GetUsersHandler) Handle(c *fiber.Ctx) error {
	users, err := h.repo.GetUsers(c.Context())
	if err != nil {
		zap.L().Error("error while getting users", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	users = lo.Filter(users, func(u entity.User, _ int) bool {
		return u.Role == utils.ROLE_USER
	})

	userResponses := lo.Map(users, func(u entity.User, _ int) response.UserResponse {
		return response.UserResponse{
			UserId:      u.ID,
			Email:       u.Email,
			FirstName:   u.FirstName,
			LastName:    u.LastName,
			Gender:      u.Gender,
			DateOfBirth: u.DateOfBirth,
			Role:        u.Role,
		}
	})

	return c.JSON(response.UserListResponse{Users: userResponses})
}
