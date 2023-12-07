// auth/jwt.go
package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

//TODO:: generate a proper secret key and store in the secrets.json and read from there
const secretKey = "my_super_secure_secret_key_123!@#"

//TODO:: add role to the claims
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

//TODO:: if you are using middleware in this package which is not recommended, you should make this function private
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

//TODO:: middlewares should be in another package called middlewares
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
