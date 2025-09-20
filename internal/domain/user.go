package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName    string `json:"full_name" gorm:"not null"`
	Email       string `gorm:"unique" json:"email"`
	IsConfirmed bool   `json:"is_confirmed" gorm:"default:false"`
	Password    string `json:"-"`
}
