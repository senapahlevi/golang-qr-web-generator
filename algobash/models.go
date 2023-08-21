package algobash

type Customer struct {
	ID        int    `gorm:"id" json:"id"`
	FirstName string `gorm:"firstname" json:"firstname"`
	LastName  string `gorm:"lastname" json:"lastname"`
}

func (Customer) TableName() string {
	return "customer"
}
