package helpers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chandan782/Pismo/api/helpers"
)

func TestGetAmount(t *testing.T) {
	tests := []struct {
		name            string
		amount          float64
		operationTypeId int
		expectedResult  float64
	}{
		{
			name:            "Normal Purchase",
			amount:          100.0,
			operationTypeId: helpers.NORMAL_PURCHASE_ID,
			expectedResult:  -100.0,
		},
		{
			name:            "Purchase with Installments",
			amount:          300.0,
			operationTypeId: helpers.PURCHASE_WITH_INSTALLMENT_ID,
			expectedResult:  -100.0,
		},
		{
			name:            "Withdrawal",
			amount:          50.0,
			operationTypeId: helpers.WITHDRAWAL_ID,
			expectedResult:  -50.0,
		},
		{
			name:            "Credit Voucher",
			amount:          200.0,
			operationTypeId: helpers.CREDIT_VOUCHER_ID,
			expectedResult:  200.0,
		},
		{
			name:            "Invalid Operation Type",
			amount:          100.0,
			operationTypeId: 999,
			expectedResult:  0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := helpers.GetAmount(tt.amount, tt.operationTypeId)
			assert.Equal(t, tt.expectedResult, result)
		})
	}
}
