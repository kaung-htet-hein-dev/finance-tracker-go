package domain

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Amount float64 `json:"amount" gorm:"not null,unique"`
	Note   string  `json:"note"`
}
