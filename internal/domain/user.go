package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID           uint          `gorm:"primaryKey" json:"id"`
	FullName     string        `json:"full_name"`
	Email        string        `gorm:"unique" json:"email"`
	Password     string        `json:"-"`
	IsConfirmed  bool          `json:"is_confirmed" gorm:"default:false"`
	Categories   []Category    `gorm:"foreignKey:UserID" json:"categories,omitempty"`
	Transactions []Transaction `gorm:"foreignKey:UserID" json:"transactions,omitempty"`
}
