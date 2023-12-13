package validators

import (
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateCreateNewUserRequest(t *testing.T) {
	t.Run("nil req should return error", func(t *testing.T) {
		err := ValidateCreateNewUserRequest(nil)

		assert.EqualError(t, err, ErrRequestBodyCannotBeEmpty.Error())
	})

	t.Run("empty email should return error", func(t *testing.T) {
		err := ValidateCreateNewUserRequest(&request.NewUserRequest{
			Email:     "",
			Password:  "password",
			Firstname: "firstname",
			LastName:  "lastname",
		})

		errorContains(t, err, "email cannot be empty")
	})

	t.Run("empty password should return error", func(t *testing.T) {
		err := ValidateCreateNewUserRequest(&request.NewUserRequest{
			Email:     "email",
			Password:  "",
			Firstname: "firstname",
			LastName:  "lastname",
		})

		errorContains(t, err, "password cannot be empty")
	})

	t.Run("empty firstname should return error", func(t *testing.T) {
		err := ValidateCreateNewUserRequest(&request.NewUserRequest{
			Email:     "email",
			Password:  "password",
			Firstname: "",
			LastName:  "lastname",
		})

		errorContains(t, err, "first name cannot be empty")
	})

	t.Run("empty lastname should return error", func(t *testing.T) {
		err := ValidateCreateNewUserRequest(&request.NewUserRequest{
			Email:     "email",
			Password:  "password",
			Firstname: "firstname",
			LastName:  "",
		})
		errorContains(t, err, "last name cannot be empty")
	})

	t.Run("empty role should return error", func(t *testing.T) {
		err := ValidateCreateNewUserRequest(&request.NewUserRequest{
			Email:     "email",
			Password:  "password",
			Firstname: "firstname",
			LastName:  "lastname",
			Role:      "",
		})
		errorContains(t, err, "role cannot be empty")
	})

	t.Run("invalid role should return error", func(t *testing.T) {
		err := ValidateCreateNewUserRequest(&request.NewUserRequest{
			Email:     "email",
			Password:  "password",
			Firstname: "firstname",
			LastName:  "lastname",
			Role:      "invalid_role",
		})
		errorContains(t, err, "role can be either admin or user")
	})

}

func errorContains(t *testing.T, err error, errorMessages ...string) {
	for _, e := range errorMessages {
		assert.ErrorContains(t, err, e)
	}
}
