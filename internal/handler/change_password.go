package handler

import (
	"errors"
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type ChangePasswordHandler struct {
	repo database.Repository
}

func NewChangePasswordHandler(repo database.Repository) *ChangePasswordHandler {
	return &ChangePasswordHandler{
		repo: repo,
	}
}

func (h *ChangePasswordHandler) Handle(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	req := new(request.ChangePassword)
	if err := c.BodyParser(req); err != nil {
		zap.L().Error("error while parsing request body", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	if err := validators.Validate(req); err != nil {
		zap.L().Error("error while validating request body", zap.Error(err))
		c.Status(fiber.StatusBadRequest)
		return err
	}

	user, err := h.repo.GetUserById(c.Context(), userId)
	if err != nil {
		zap.L().Error("error while getting user", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	if user == nil {
		zap.L().Error("user not found")
		return c.Status(fiber.StatusNotFound).SendString("user not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword)); err != nil {
		zap.L().Error("invalid credentials", zap.Error(err))
		c.Status(fiber.StatusBadRequest)
		return errors.New("invalid credentials")
	}

	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), 10)
	if err != nil {
		zap.L().Error("error while encrypting the password", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return errors.New("error while encrypting the password")
	}

	if err = h.repo.UpdatePassword(c.Context(), string(hashedNewPassword), *user); err != nil {
		zap.L().Error("error while updating the password", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.Status(fiber.StatusOK).SendString("password updated successfully")
}
