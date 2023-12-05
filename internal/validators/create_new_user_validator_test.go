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
		err := ValidateCreateNewUserRequest(&request.CreateNewUserRequest{
			Email:     "",
			Password:  "password",
			Firstname: "firstname",
			LastName:  "lastname",
		})

		errorContains(t, err, "email cannot be empty")
	})

	//TODO:: complete missing test cases

}

func errorContains(t *testing.T, err error, errorMessages ...string) {
	for _, e := range errorMessages {
		assert.ErrorContains(t, err, e)
	}
}
