package models

type RolesUser struct {
	ID   int    `gorm:"id" json:"id"`
	Name string `gorm:"name" json:"name"`
}

func (RolesUser) TableName() string {
	return "roles_user"
}
