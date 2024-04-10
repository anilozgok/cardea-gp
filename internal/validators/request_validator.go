package validators

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func Validate(r interface{}) error {
	var errToReturn error

	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			errToReturn = nil
		}

		for _, err := range err.(validator.ValidationErrors) {
			errToReturn = errors.New(fmt.Sprintf("%s %s %s %s ", err.Namespace(), err.Field(), err.Tag(), err.Type()))
			break
		}
	}

	return errToReturn
}
