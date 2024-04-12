package schemas

type Transaction struct {
	ID              string  `gorm:"id"`
	AccountID       string  `gorm:"accountid"`
	OperationTypeID int     `gorm:"operationtypeid"`
	Amount          float64 `gorm:"amount"`
	EventDate       string  `gorm:"eventdate"`
}
