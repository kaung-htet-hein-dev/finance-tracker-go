package usecase

import (
	"errors"
	"kaung-htet-hein-dev/finance-tracker-go/internal/domain"
	"kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/repository"
)

type TransactionUsecase struct {
	repo         *repository.TransactionRepository
	categoryRepo *repository.CategoryRepository
}

func NewTransactionUsecase(
	repo *repository.TransactionRepository,
	categoryRepo *repository.CategoryRepository,
) *TransactionUsecase {
	return &TransactionUsecase{
		repo:         repo,
		categoryRepo: categoryRepo,
	}
}

func (u *TransactionUsecase) CreateTransaction(
	amount float64,
	note string,
	transactionType string,
	categoryID uint,
	userID uint,
) (*domain.Transaction, error) {

	// Check if category exists and belongs to the user
	category, err := u.categoryRepo.FindByID(categoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	if category.UserID != userID {
		return nil, errors.New("category does not belong to this user")
	}

	transaction := &domain.Transaction{
		Amount:     amount,
		Note:       note,
		Type:       transactionType,
		CategoryID: categoryID,
		UserID:     userID,
	}

	if err := u.repo.Create(transaction); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (u *TransactionUsecase) GetTransactions(userID uint) ([]domain.Transaction, error) {
	return u.repo.FindByUserID(userID)
}

func (u *TransactionUsecase) GetTransactionByID(id uint, userID uint) (*domain.Transaction, error) {
	transaction, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Ensure the transaction belongs to the user
	if transaction.UserID != userID {
		return nil, errors.New("transaction not found")
	}

	return transaction, nil
}

func (u *TransactionUsecase) FilterTransactions(
	userID uint,
	transactionType string,
	categoryID uint,
) ([]domain.Transaction, error) {
	return u.repo.FindWithFilters(userID, transactionType, categoryID)
}

func (u *TransactionUsecase) UpdateTransaction(
	id uint,
	amount float64,
	note string,
	transactionType string,
	categoryID uint,
	userID uint,
) (*domain.Transaction, error) {
	// Get the transaction and ensure it belongs to the user
	transaction, err := u.GetTransactionByID(id, userID)
	if err != nil {
		return nil, err
	}

	// Validate transaction type
	if transactionType != "income" && transactionType != "expense" {
		return nil, errors.New("invalid transaction type: must be 'income' or 'expense'")
	}

	// Check if category exists and belongs to the user
	category, err := u.categoryRepo.FindByID(categoryID)
	if err != nil {
		return nil, errors.New("category not found")
	}

	if category.UserID != userID {
		return nil, errors.New("category does not belong to this user")
	}

	// Update fields
	transaction.Amount = amount
	transaction.Note = note
	transaction.Type = transactionType
	transaction.CategoryID = categoryID

	if err := u.repo.Update(transaction); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (u *TransactionUsecase) DeleteTransaction(id uint, userID uint) error {
	// Get the transaction and ensure it belongs to the user
	transaction, err := u.GetTransactionByID(id, userID)
	if err != nil {
		return err
	}

	return u.repo.Delete(transaction.ID)
}
