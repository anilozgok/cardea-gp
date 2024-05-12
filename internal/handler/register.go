package handler

import (
	"errors"
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/entity"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type RegisterHandler struct {
	repo database.Repository
}

func NewRegisterHandler(repo database.Repository) *RegisterHandler {
	return &RegisterHandler{
		repo: repo,
	}
}

func (h *RegisterHandler) Handle(c *fiber.Ctx) error {
	req := new(request.NewUser)
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

	if maybeUser != nil {
		zap.L().Error("user already exists", zap.Error(err))
		c.Status(fiber.StatusBadRequest)
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		zap.L().Error("error while encrypting the password", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return errors.New("error while encrypting the password")
	}

	user := entity.User{
		FirstName:   req.Firstname,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    string(hashedPassword),
		Gender:      req.Gender,
		DateOfBirth: req.DateOfBirth,
		Role:        strings.ToLower(req.Role),
	}

	if err = h.repo.CreateNewUser(c.Context(), &user); err != nil {
		zap.L().Error("error while creating new user", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
