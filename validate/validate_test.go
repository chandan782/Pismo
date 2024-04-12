package validate_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/chandan782/Pismo/validate"
)

type TestStruct struct {
	Field1 string `json:"field1" validate:"required"`
	Field2 int    `json:"field2" validate:"gte=10"`
}

func TestValidateStructs(t *testing.T) {
	// create a new instance of the validator
	v := validate.New()

	tests := []struct {
		name     string
		req      interface{}
		expected error
	}{
		{
			name:     "Valid struct",
			req:      TestStruct{Field1: "value", Field2: 15},
			expected: nil,
		},
		{
			name:     "Missing required fields",
			req:      TestStruct{Field2: 15},
			expected: errors.New("Field 'Field1' failed validation with tag 'required'"),
		},
		{
			name:     "Invalid integer fields",
			req:      TestStruct{Field1: "value", Field2: 5},
			expected: errors.New("Field 'Field2' failed validation with tag 'gte'"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.ValidateStructs(tt.req)

			if err == nil && tt.expected != nil {
				t.Errorf("expected error, got nil")
			} else if err != nil && !strings.Contains(err.Error(), tt.expected.Error()) {
				t.Errorf("validation error mismatch. expected substring: %v, got: %v", tt.expected, err.Error())
			}
		})
	}
}
