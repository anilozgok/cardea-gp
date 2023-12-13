package validators

import (
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"testing"
)

func TestValidateLoginRequest(t *testing.T) {
	t.Run("nil req should return error", func(t *testing.T) {
		err := ValidateLoginRequest(nil)

		errorContains(t, err, ErrRequestBodyCannotBeEmpty.Error())
	})

	t.Run("empty email should return error", func(t *testing.T) {
		err := ValidateLoginRequest(&request.LoginRequest{
			Email:    "",
			Password: "password",
		})

		errorContains(t, err, "email cannot be empty")
	})

	t.Run("empty password should return error", func(t *testing.T) {
		err := ValidateLoginRequest(&request.LoginRequest{
			Email:    "email",
			Password: "",
		})

		errorContains(t, err, "password cannot be empty")
	})
}
