package validators

import (
	"errors"
	"github.com/anilozgok/cardea-gp/internal/model/request"
	"github.com/anilozgok/cardea-gp/pkg/utils"
)

var (
	ErrRequestBodyCannotBeEmpty = errors.New("request body cannot be empty")
)

func ValidateNewUserRequest(req *request.NewUserRequest) error {
	if req == nil {
		return ErrRequestBodyCannotBeEmpty
	}

	var err error
	if req.Firstname == "" {
		err = errors.Join(err, errors.New("first name cannot be empty"))
	}

	if req.LastName == "" {
		err = errors.Join(err, errors.New("last name cannot be empty"))
	}

	if req.DateOfBirth.IsZero() {
		err = errors.Join(err, errors.New("date of birth cannot be empty"))
	}

	if req.Gender == "" {
		err = errors.Join(err, errors.New("gender cannot be empty"))
	}

	if req.Height == 0 {
		err = errors.Join(err, errors.New("height cannot be empty"))
	}

	if req.Weight == 0 {
		err = errors.Join(err, errors.New("weight cannot be empty"))
	}

	if req.Email == "" {
		err = errors.Join(err, errors.New("email cannot be empty"))
	}

	if req.Password == "" {
		err = errors.Join(err, errors.New("password cannot be empty"))
	}

	if req.Role == "" {
		err = errors.Join(err, errors.New("role cannot be empty"))
	}

	if req.Role != utils.ROLE_ADMIN && req.Role != utils.ROLE_USER || req.Role != utils.ROLE_COACH {
		err = errors.Join(err, errors.New("invalid role"))
	}

	return err
}
