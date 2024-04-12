package models

import "github.com/go-playground/validator/v10"

// CreateTransactionRequest represents the request to create a transaction
type CreateTransactionRequest struct {
	AccountID       string  `json:"account_id" validate:"required"`
	OperationTypeID int     `json:"operation_type_id" validate:"required,oneof=1 2 3 4"`
	Amount          float64 `json:"amount" validate:"required,gt=0"`
}

type CreateTransactionResponse struct {
	ID string `json:"id"`
}

// validate validates the struct fields using the validator package
func (req *CreateTransactionRequest) Validate() error {
	v := validator.New()
	if err := v.Struct(req); err != nil {
		return formatValidationErrors(err.(validator.ValidationErrors))
	}
	return nil
}
