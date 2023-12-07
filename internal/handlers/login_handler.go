package handlers

import (
	"errors"
	"github.com/anilozgok/cardea-gp/internal/auth"
	"github.com/anilozgok/cardea-gp/internal/entities"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/internal/validators"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginHandler(db *gorm.DB) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		req := new(request.LoginRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}
		if err := validators.ValidateLoginRequest(req); err != nil {
			return err
		}

		user := entities.User{}
		if result := db.Where("email = ?", req.Email).First(&user); result.Error != nil {
			return result.Error
		}
		if result := db.Where(&user).First(&user); result.Error != nil {
			return result.Error
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			return errors.New("invalid credentials")
		}
		tokenString, err := auth.CreateToken(req.Email)
		if err != nil {
			return err
		}

		c.Cookie(&fiber.Cookie{
			Name:  "jwt",
			Value: tokenString,
		})

		return c.JSON(map[string]string{"token": tokenString})
	}

}
