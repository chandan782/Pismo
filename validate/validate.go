package validate

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validate interface {
	ValidateStructs(req interface{}) error
}

type validate struct {
	valid *validator.Validate
}

func New() Validate {
	v := validator.New()
	return &validate{
		valid: v,
	}
}

// validate validates the struct fields using the validator package
func (v *validate) ValidateStructs(req interface{}) error {
	if err := v.valid.Struct(req); err != nil {
		return v.formatValidationErrors(err.(validator.ValidationErrors))
	}
	return nil
}

// formatValidationErrors formats validation errors into a readable string
func (v *validate) formatValidationErrors(err validator.ValidationErrors) error {
	var errMsgs []string
	for _, e := range err {
		field := e.Field()
		tag := e.Tag()
		msg := fmt.Sprintf("Field '%s' failed validation with tag '%s'", field, tag)
		errMsgs = append(errMsgs, msg)
	}
	return errors.New(strings.Join(errMsgs, ", "))
}
