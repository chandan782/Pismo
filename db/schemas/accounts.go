package schemas

type Account struct {
	ID             string `gorm:"id"`
	DocumentNumber string `gorm:"documentnumber"`
}
