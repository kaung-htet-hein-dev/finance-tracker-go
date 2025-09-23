package repository

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/domain"
	"kaung-htet-hein-dev/finance-tracker-go/pkg"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) Create(transaction *domain.Transaction) error {
	err := r.DB.Create(transaction).Error

	if err != nil {
		return pkg.HandleGormError(err, "transaction")
	}

	return nil
}

func (r *TransactionRepository) FindAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	result := r.DB.Find(&transactions)
	return transactions, result.Error
}

func (r *TransactionRepository) FindByID(id uint) (*domain.Transaction, error) {
	var transaction domain.Transaction
	result := r.DB.First(&transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &transaction, nil
}

func (r *TransactionRepository) FindByUserID(userID uint) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	result := r.DB.Where("user_id = ?", userID).Find(&transactions)
	return transactions, result.Error
}

func (r *TransactionRepository) Update(transaction *domain.Transaction) error {
	return r.DB.Save(transaction).Error
}

func (r *TransactionRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.Transaction{}, id).Error
}

// Optional: Add more complex queries
func (r *TransactionRepository) FindWithFilters(userID uint, transactionType string, categoryID uint) ([]domain.Transaction, error) {
	query := r.DB.Where("user_id = ?", userID)

	if transactionType != "" {
		query = query.Where("type = ?", transactionType)
	}

	if categoryID > 0 {
		query = query.Where("category_id = ?", categoryID)
	}

	var transactions []domain.Transaction
	result := query.Find(&transactions)
	return transactions, result.Error
}
