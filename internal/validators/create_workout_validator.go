package validators

import (
	"errors"
	"github.com/anilozgok/cardea-gp/internal/model/request"
)

func ValidateCreateWorkoutRequest(req *request.CreateWorkoutRequest) error {
	if req == nil {
		return ErrRequestBodyCannotBeEmpty
	}

	if req.Name == "" {
		return errors.New("name of the workout cannot be empty")
	}

	if req.Description == "" {
		return errors.New("description of the workout cannot be empty")
	}

	if req.Area == "" {
		return errors.New("area of the workout cannot be empty")
	}

	if req.Rep == 0 {
		return errors.New("rep of the workout cannot be empty")
	}

	if req.Sets == 0 {
		return errors.New("sets of the workout cannot be empty")
	}

	return nil
}
