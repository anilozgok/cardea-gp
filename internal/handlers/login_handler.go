package handlers

import (
	"errors"
	"github.com/anilozgok/cardea-gp/internal/jwt"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/repository"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginHandler struct {
	repo repository.Repository
}

func NewLoginHandler(repo repository.Repository) *LoginHandler {
	return &LoginHandler{repo: repo}
}

func (h *LoginHandler) Handle(c *fiber.Ctx) error {
	req := new(request.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	if err := validators.ValidateLoginRequest(req); err != nil {
		return err
	}

	user, err := h.repo.GetUserByEmail(c.Context(), req.Email)
	if err != nil {
		return err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.New("invalid credentials")
	}

	opts := jwt.Opts{
		UserId: uint32(user.ID),
		Email:  user.Email,
		Role:   user.Role,
	}

	token, err := jwt.CreateToken(opts)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:    "authSession",
		Value:   token,
		Expires: time.Now().Add(6 * time.Hour), // token is valid for 6 hours
	})

	return c.Status(fiber.StatusOK).SendString(token)
}
