package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"uniqueIndex;not null" validate:"required,min=3,max=50"`
	Email     string         `json:"email" gorm:"uniqueIndex;not null" validate:"required,email"`
	Password  string         `json:"-" gorm:"not null" validate:"required,min=6"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)

type Transaction struct {
	ID          uint            `json:"id" gorm:"primaryKey"`
	UserID      uint            `json:"user_id" gorm:"not null" validate:"required"`
	User        User            `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Type        TransactionType `json:"type" gorm:"not null" validate:"required,oneof=income expense"`
	Amount      float64         `json:"amount" gorm:"not null" validate:"required,gt=0"`
	Category    string          `json:"category" gorm:"not null" validate:"required,min=1,max=100"`
	Description string          `json:"description" validate:"max=500"`
	Date        time.Time       `json:"date" gorm:"not null" validate:"required"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   gorm.DeletedAt  `json:"-" gorm:"index"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type TransactionRequest struct {
	Type        TransactionType `json:"type" validate:"required,oneof=income expense"`
	Amount      float64         `json:"amount" validate:"required,gt=0"`
	Category    string          `json:"category" validate:"required,min=1,max=100"`
	Description string          `json:"description" validate:"max=500"`
	Date        time.Time       `json:"date" validate:"required"`
}

type InsightResponse struct {
	TotalIncome       float64            `json:"total_income"`
	TotalExpenses     float64            `json:"total_expenses"`
	NetIncome         float64            `json:"net_income"`
	TopCategories     []CategoryInsight  `json:"top_categories"`
	MonthlyTrend      []MonthlyTrend     `json:"monthly_trend"`
	Recommendations   []string           `json:"recommendations"`
	SavingsRate       float64            `json:"savings_rate"`
	AverageExpense    float64            `json:"average_expense"`
}

type CategoryInsight struct {
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
	Count    int     `json:"count"`
}

type MonthlyTrend struct {
	Month    string  `json:"month"`
	Income   float64 `json:"income"`
	Expenses float64 `json:"expenses"`
}