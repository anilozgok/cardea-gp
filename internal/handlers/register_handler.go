package handlers

import (
	"errors"
	"github.com/anilozgok/cardea-gp/internal/entities"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/repository"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type RegisterHandler struct {
	repo repository.Repository
}

func NewRegisterHandler(repo repository.Repository) *RegisterHandler {
	return &RegisterHandler{
		repo: repo,
	}
}

func (h *RegisterHandler) Handle(c *fiber.Ctx) error {
	req := new(request.NewUserRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	if err := validators.ValidateCreateNewUserRequest(req); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return errors.New("error while encrypting the password")
	}

	user := entities.User{
		Email:     req.Email,
		Password:  string(hashedPassword),
		Firstname: req.Firstname,
		LastName:  req.LastName,
		Role:      strings.ToLower(req.Role),
	}

	if err = h.repo.CreateNewUser(c.Context(), &user); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}
