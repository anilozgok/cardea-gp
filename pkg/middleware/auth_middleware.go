package middleware

import (
	"errors"
	"github.com/anilozgok/cardea-gp/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("authSession")

	if token == "" {
		zap.L().Error("no authorization token found")
		return errors.New("no authorization token found")
	}

	claims, err := jwt.VerifyToken(token)
	if err != nil {
		zap.L().Error("invalid jwt token", zap.Error(err), zap.String("token", token))
		return err
	}

	if claims == nil || claims.Email == "" || claims.UserId == 0 {
		zap.L().Error("no claim found", zap.String("token", token))
		return errors.New("no claim found")
	}

	c.Locals("user", claims)
	c.Locals("role", claims.Role)
	c.Locals("email", claims.Email)
	c.Locals("userId", claims.UserId)

	return c.Next()
}
