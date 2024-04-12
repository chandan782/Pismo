package helpers

const (
	// defined all transactions key
	NORMAL_PURCHASE_ID           = 1
	PURCHASE_WITH_INSTALLMENT_ID = 2
	WITHDRAWAL_ID                = 3
	CREDIT_VOUCHER_ID            = 4

	// defined all transactions value
	NORMAL_PURCHASE_VALUE           = "Normal Purchase"
	PURCHASE_WITH_INSTALLMENT_VALUE = "Purchase with Installments"
	WITHDRAWAL_VALUE                = "Withdrawal"
	CREDIT_VOUCHER_VALUE            = "Credit Voucher"
)

func GetAmount(amount float64, operationTypeId int) float64 {
	switch operationTypeId {
	case NORMAL_PURCHASE_ID:
		return amount * -1
	case PURCHASE_WITH_INSTALLMENT_ID:
		return -(amount / 3) // assuming we have 3 months installment option
	case WITHDRAWAL_ID:
		return amount * -1
	case CREDIT_VOUCHER_ID:
		return amount
	}

	return 0.0
}
