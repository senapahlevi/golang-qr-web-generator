package models

import "time"

type QRcode struct {
	ID        int       `gorm:"id" json:"id"`
	Text      string    `gorm:"text" json:"text"`
	TextURL   string    `gorm:"text_url" json:"text_url"`
	CreatedAt time.Time `gorm:"updated_at" json:"updated_at"`
	UpdatedAt time.Time `gorm:"created_at" json:"created_at"`
}

func (QRcode) TableName() string {
	return "QR_code"
}
