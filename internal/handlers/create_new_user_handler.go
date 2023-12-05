package handlers

import (
	"errors"
	"github.com/anilozgok/cardea-gp/internal/entities"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
)

func CreateNewUserHandler(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		req := new(request.CreateNewUserRequest)
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

		if result := db.Save(&user); result.Error != nil {
			return result.Error
		}

		//TODO:: return jwt
		return c.JSON("Ok")
	}
}
