package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"unique;not null" json:"name"`
	UserID uint   `json:"user_id"`
	User   *User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}
