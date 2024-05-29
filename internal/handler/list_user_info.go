package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/response"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ListUserInfoHandler struct {
	repo database.Repository
}

func NewLListUserInfoHandler(repo database.Repository) *ListUserInfoHandler {
	return &ListUserInfoHandler{repo: repo}
}

func (h *ListUserInfoHandler) Handle(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	userInfo, err := h.repo.GetUserById(c.Context(), userId)
	if err != nil {
		zap.L().Error("error while fetching user information", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}
	userResponse := response.UserResponse{
		UserId:      userInfo.ID,
		Email:       userInfo.Email,
		FirstName:   userInfo.FirstName,
		LastName:    userInfo.LastName,
		Gender:      userInfo.Gender,
		DateOfBirth: userInfo.DateOfBirth,
		Role:        userInfo.Role,
	}

	return c.JSON(userResponse)
}
