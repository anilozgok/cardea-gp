package validators

import (
	"errors"
	"github.com/anilozgok/cardea-gp/internal/model/request"
)

func ValidateLoginRequest(req *request.LoginRequest) error {
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

	return err
}
