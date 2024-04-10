package handler

import (
	"errors"
	"github.com/anilozgok/cardea-gp/internal/database"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/anilozgok/cardea-gp/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginHandler struct {
	repo database.Repository
}

func NewLoginHandler(repo database.Repository) *LoginHandler {
	return &LoginHandler{repo: repo}
}

func (h *LoginHandler) Handle(c *fiber.Ctx) error {
	req := new(request.LoginRequest)
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

	user, err := h.repo.GetUserByEmail(c.Context(), req.Email)
	if err != nil {
		zap.L().Error("error while getting user by email", zap.Error(err))
		return err
	}

	if user == nil {
		zap.L().Error("user not found", zap.String("email", req.Email))
		c.Status(fiber.StatusNotFound)
		return errors.New("user not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		zap.L().Error("invalid credentials", zap.Error(err))
		c.Status(fiber.StatusBadRequest)
		return errors.New("invalid credentials")
	}

	opts := jwt.Opts{
		UserId: uint32(user.ID),
		Email:  user.Email,
		Role:   user.Role,
	}

	token, err := jwt.CreateToken(opts)
	if err != nil {
		zap.L().Error("error while creating jwt token", zap.Error(err))
		c.Status(fiber.StatusInternalServerError)
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:    "authSession",
		Value:   token,
		Expires: time.Now().Add(6 * time.Hour), // token is valid for 6 hours
	})

	return c.Status(fiber.StatusOK).SendString(token)
}
