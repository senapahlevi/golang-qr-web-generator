package login

type LoginRequest struct {
	Username string `gorm:"username" `
	Password string `gorm:"password"`
}
