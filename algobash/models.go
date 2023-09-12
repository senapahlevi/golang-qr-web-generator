package algobash

type Customer struct {
	ID        int    `gorm:"id" json:"id"`
	FirstName string `gorm:"first_name" json:"firstname"`
	LastName  string `gorm:"last_name" json:"lastname"`
}

func (Customer) TableName() string {
	return "customer"
}
