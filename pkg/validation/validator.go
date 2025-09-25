package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// Validator wrapper
type Validator struct {
	Validator *validator.Validate
}

// Validator wrapper constructor
func New() *Validator {
	return &Validator{
		Validator: validator.New(),
	}
}

// Validate object with custom error messages
func (v *Validator) Validate(obj interface{}) []string {
	if errs := v.Validator.Struct(obj); errs != nil {
		errMsgs := make([]string, 0)

		for _, err := range errs.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: %s",
				err.Field(),
				err.Tag(),
			))
		}

		return errMsgs
	}

	return nil
}
