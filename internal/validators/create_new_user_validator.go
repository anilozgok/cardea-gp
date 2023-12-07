package validators

import (
	"errors"
	"github.com/anilozgok/cardea-gp/internal/model/request"
)

var (
	ErrRequestBodyCannotBeEmpty = errors.New("request body cannot be empty")
)

func ValidateCreateNewUserRequest(req *request.CreateNewUserRequest) error {
	if req == nil {
		return ErrRequestBodyCannotBeEmpty
	}

	var err error
	if req.Email == "" {
		err = errors.Join(err, errors.New("email cannot be empty"))
	}

	if req.Password == "" {
		err = errors.Join(err, errors.New("password cannot be empty"))
	}

	if req.Firstname == "" {
		err = errors.Join(err, errors.New("first name cannot be empty"))
	}

	if req.LastName == "" {
		err = errors.Join(err, errors.New("last name cannot be empty"))
	}

	if req.Role == "" {
		err = errors.Join(err, errors.New("role cannot be empty"))
	}
	if req.Role != "admin" && req.Role != "user" {
		err = errors.Join(err, errors.New("role can be either admin or user"))
	}

	return err
}
