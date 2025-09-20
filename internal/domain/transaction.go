package domain

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID         uint      `gorm:"primaryKey" json:"id"`
	Amount     float64   `json:"amount"`
	Note       string    `json:"note"`
	Type       string    `json:"type"`
	UserID     uint      `json:"user_id"`
	User       *User     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	CategoryID uint      `json:"category_id"`
	Category   *Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}
