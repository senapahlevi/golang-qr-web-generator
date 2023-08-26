package models

type Customer struct {
	ID        int       `gorm:"id" json:"id"`
	FirstName string    `gorm:"firstname" json:"firstname"`
	LastName  string    `gorm:"lastname" json:"lastname"`
	RoleID    string    `gorm:"role_id" json:"role_id"`
	Email     string    `gorm:"email"`
	Username  string    `gorm:"username" `
	Password  string    `gorm:"password"`
	Token     string    `gorm:"token" `
	Customer  RolesUser `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Customer) TableName() string {
	return "customer"
}
