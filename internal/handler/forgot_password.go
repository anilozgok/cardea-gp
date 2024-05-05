package handler

import (
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UpdatePasswordHandler struct {
	repo database.Repository
}

func NewUpdatePasswordHandler(repo database.Repository) *UpdatePasswordHandler {
	return &UpdatePasswordHandler{
		repo: repo,
	}
}

func (h *UpdatePasswordHandler) Handle(c *fiber.Ctx) error {
	req := new(request.ForgotPassword)
	if err := c.BodyParser(req); err != nil {
		zap.L().Error("error while parsing request body", zap.Error(err))
		c.Status(fiber.StatusBadRequest)
		return err
	}

	if err := validators.Validate(req); err != nil {
		zap.L().Error("error while validating request body", zap.Error(err))
		c.Status(fiber.StatusBadRequest)
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		zap.L().Error("error while encrypting the password", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	user := c.Locals("user").(entity.User)

	if err = h.repo.UpdatePassword(c.Context(), string(hashedPassword), user); err != nil {
		zap.L().Error("error while updating the password", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
