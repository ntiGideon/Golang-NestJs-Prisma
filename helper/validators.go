package helper

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

func RequestValidators(post interface{}) error {
	validate := validator.New()
	err := validate.Struct(post)
	if err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		return errs
	}
	return nil
}
