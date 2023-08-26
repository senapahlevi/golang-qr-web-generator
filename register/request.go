package register

type RegisterRequest struct {
	FirstName string `gorm:"firstname" json:"firstname"`
	LastName  string `gorm:"lastname" json:"lastname"`
	Email     string `gorm:"email"`
	Username  string `gorm:"username" `
	Password  string `gorm:"password"`
}
