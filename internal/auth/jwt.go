// auth/jwt.go
package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "my_super_secure_secret_key_123!@#"

// It's best to store this in an environment variable

// Claims represents the claims that can be encoded in a JWT
type Claims struct {
	Email string `json:"email"`
	jwt.MapClaims
}

// CreateToken creates a JWT token with the given email and expiration time.
func CreateToken(email string) (string, error) {
	claims := Claims{
		Email: email,
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	println(tokenString)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerifyToken verifies the JWT token and returns the claims
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token, nil
}
func JWTMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")

		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "unauthorized: token cannot be empty",
			})
		}

		token, err := VerifyToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"unauthorized": "invalid token",
			})
		}

		// Extract claims from token
		claims, ok := token.Claims.(*Claims)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "internal server error: unable to extract token claims",
			})
		}

		c.Locals("claims", claims)
		return c.Next()
	}
}
