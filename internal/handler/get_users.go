package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/entities"
	"github.com/anilozgok/cardea-gp/internal/model/response"
	"github.com/gofiber/fiber/v2"
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

	res := mapToResponse(users)

	return c.JSON(res)
}

func mapToResponse(users []entities.User) response.UserListResponse {
	var userResponses []response.UserResponse
	for _, u := range users {
		userResponses = append(userResponses, response.UserResponse{
			UserId:      uint32(u.ID),
			Email:       u.Email,
			FirstName:   u.FirstName,
			LastName:    u.LastName,
			Gender:      u.Gender,
			Height:      u.Height,
			Weight:      u.Weight,
			DateOfBirth: u.DateOfBirth,
			Role:        u.Role,
		})
	}
	return response.UserListResponse{Users: userResponses}
}
