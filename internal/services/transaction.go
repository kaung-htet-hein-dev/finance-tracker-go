package services

import (
	"errors"
	"finance-tracker-go/internal/database"
	"finance-tracker-go/internal/models"

	"gorm.io/gorm"
)

type TransactionService struct{}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (s *TransactionService) CreateTransaction(userID uint, req models.TransactionRequest) (*models.Transaction, error) {
	transaction := models.Transaction{
		UserID:      userID,
		Type:        req.Type,
		Amount:      req.Amount,
		Category:    req.Category,
		Description: req.Description,
		Date:        req.Date,
	}

	if err := database.GetDB().Create(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (s *TransactionService) GetTransactions(userID uint, limit, offset int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := database.GetDB().Where("user_id = ?", userID).
		Order("date DESC").
		Limit(limit).
		Offset(offset)

	if err := query.Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}

func (s *TransactionService) GetTransactionByID(userID, transactionID uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := database.GetDB().Where("id = ? AND user_id = ?", transactionID, userID).First(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("transaction not found")
		}
		return nil, err
	}

	return &transaction, nil
}

func (s *TransactionService) UpdateTransaction(userID, transactionID uint, req models.TransactionRequest) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := database.GetDB().Where("id = ? AND user_id = ?", transactionID, userID).First(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("transaction not found")
		}
		return nil, err
	}

	transaction.Type = req.Type
	transaction.Amount = req.Amount
	transaction.Category = req.Category
	transaction.Description = req.Description
	transaction.Date = req.Date

	if err := database.GetDB().Save(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (s *TransactionService) DeleteTransaction(userID, transactionID uint) error {
	result := database.GetDB().Where("id = ? AND user_id = ?", transactionID, userID).Delete(&models.Transaction{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("transaction not found")
	}

	return nil
}