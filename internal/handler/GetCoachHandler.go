package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/response"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type GetCoachHandler struct {
	repo database.Repository
}

func NewGetCoachHandler(repo database.Repository) *GetCoachHandler {
	return &GetCoachHandler{repo: repo}
}

func (h *GetCoachHandler) Handle(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	coach, err := h.repo.GetCoachByUserId(c.Context(), userId)
	if err != nil {
		zap.L().Error("error while getting coach", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}

	coachResponse := response.UserResponse{
		UserId:      coach.ID,
		Email:       coach.Email,
		FirstName:   coach.FirstName,
		LastName:    coach.LastName,
		Gender:      coach.Gender,
		DateOfBirth: coach.DateOfBirth,
		Role:        coach.Role,
	}

	return c.JSON(coachResponse)
}
