package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Compare this snippet from internal/handlers/create_new_user_handler_test.go:

func TestCreateToken(t *testing.T) {
	email := "email"
	token, err := CreateToken(email)
	if err != nil {
		t.Errorf("Error while creating token: %v", err)
	}
	if token == "" {
		t.Errorf("Token is empty")
	}

}

func TestVerifyToken(t *testing.T) {
	// Generate a valid token
	email := "test@example.com"
	validToken, _ := CreateToken(email)

	// Test with a valid token
	_, err := VerifyToken(validToken)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Test with an invalid token
	_, err = VerifyToken(validToken + "invalid")
	if err == nil {
		t.Fatalf("Expected an error, got none")
	}

	// Test with an empty string
	_, err = VerifyToken("")
	if err == nil {
		t.Fatalf("Expected an error, got none")
	}
}

// ...
func TestJWTMiddleware_InvalidToken(t *testing.T) {
	app := fiber.New()
	app.Use(JWTMiddleware())

	// Mock a request with an invalid token
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")
	resp, err := app.Test(req)

	// Assert that the response status code is 401 Unauthorized
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
func TestJWTMiddleware_NoToken(t *testing.T) {
	app := fiber.New()
	app.Use(JWTMiddleware())

	// Mock a request with no token
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req)

	// Assert that the response status code is 401 Unauthorized
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)

}

//// TestJWTMiddleware
//func TestJWTMiddlewareValidToken(t *testing.T) {
//	validToken, _ := CreateToken("example@gmail.com")
//
//	app := fiber.New()
//	app.Use(JWTMiddleware())
//	// mock request
//	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
//	req.Header.Set("Authorization", "Bearer "+validToken)
//	resp, err := app.Test(req)
//	if err != nil {
//		t.Fatal(err)
//	}
//	assert.Nil(t, err)
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//
//}
