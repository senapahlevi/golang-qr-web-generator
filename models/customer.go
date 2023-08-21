package models

type Customer struct {
	ID        int       `gorm:"id" json:"id"`
	FirstName string    `gorm:"firstname" json:"firstname"`
	LastName  string    `gorm:"lastname" json:"lastname"`
	RoleID    string    `gorm:"role_id" json:"role_id"`
	Customer  RolesUser `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (Customer) TableName() string {
	return "customer"
}
