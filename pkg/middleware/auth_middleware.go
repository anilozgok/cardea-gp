package middleware

import (
	"errors"
	"github.com/anilozgok/cardea-gp/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"strings"
)

func AuthMiddleware(c *fiber.Ctx) error {
	auth := c.Get("Authorization")

	if auth == "" {
		zap.L().Error("no authorization token found")
		return errors.New("no authorization token found")
	}

	if !strings.HasPrefix(auth, "Bearer ") {
		zap.L().Error("no bearer token found", zap.String("token", auth))
		return errors.New("no bearer token found")
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	claims, err := jwt.VerifyToken(token)
	if err != nil {
		zap.L().Error("invalid jwt token", zap.Error(err), zap.String("token", auth))
		return err
	}

	if claims == nil || claims.Email == "" || claims.UserId == 0 {
		zap.L().Error("no claim found", zap.String("token", auth))
		return errors.New("no claim found")
	}

	c.Locals("user", claims)
	c.Locals("role", claims.Role)
	c.Locals("email", claims.Email)
	c.Locals("userId", claims.UserId)

	return c.Next()
}
