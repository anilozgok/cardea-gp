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

type ForgotPasswordHandler struct {
	repo database.Repository
}

func NewForgotPasswordHandler(repo database.Repository) *ForgotPasswordHandler {
	return &ForgotPasswordHandler{
		repo: repo,
	}
}

func (h *ForgotPasswordHandler) Handle(c *fiber.Ctx) error {
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

	maybeUser, err := h.repo.GetUserByEmail(c.Context(), req.Email)
	if err != nil {
		zap.L().Error("error while checking user existence", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	if maybeUser == nil {
		zap.L().Error("user not found", zap.Error(err))
		c.Status(fiber.StatusNotFound)
		return errors.New("user not found")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		zap.L().Error("error while encrypting the password", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	if err = h.repo.UpdatePassword(c.Context(), string(hashedPassword), *maybeUser); err != nil {
		zap.L().Error("error while updating the password", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
