package models

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

// CreateAccountRequest represents the request to create an account
type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number" validate:"required"`
}

type CreateAccountResponse struct {
	ID string `json:"id"`
}

// GetAccountByIdResponse represents the response for fetching an account by ID
type GetAccountByIdResponse struct {
	ID             string `json:"id"`
	DocumentNumber string `json:"document_number"`
}

// validate validates the struct fields using the validator package
func (req *CreateAccountRequest) Validate() error {
	v := validator.New()
	if err := v.Struct(req); err != nil {
		return formatValidationErrors(err.(validator.ValidationErrors))
	}
	return nil
}

// formatValidationErrors formats the validation errors into a more readable error message
func formatValidationErrors(errs validator.ValidationErrors) error {
	var errorMessages []string
	for _, e := range errs {
		// Build error message for each validation error
		errorMessage := fmt.Sprintf("%s is %s", e.Field(), e.Tag())
		errorMessages = append(errorMessages, errorMessage)
	}
	// Combine all error messages into one
	fullErrorMessage := strings.Join(errorMessages, ", ")
	return errors.Errorf(fullErrorMessage)
}
