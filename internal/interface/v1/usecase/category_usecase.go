package usecase

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/domain"
	"kaung-htet-hein-dev/finance-tracker-go/internal/infrastructure/repository"
	"kaung-htet-hein-dev/finance-tracker-go/internal/interface/v1/response"
)

type CategoryUsecase struct {
	repo *repository.CategoryRepository
}

func NewCategoryUsecase(repo *repository.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{repo: repo}
}

func (u *CategoryUsecase) CreateCategory(name string, userID uint) error {
	category := &domain.Category{Name: name, UserID: userID}
	if err := u.repo.Create(category); err != nil {
		return err
	}
	return nil
}

func (u *CategoryUsecase) GetCategories() ([]domain.Category, error) {
	return u.repo.FindAll()
}

func (u *CategoryUsecase) GetCategoryByID(id uint) (*response.CategoryResponse, error) {
	category, err := u.repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return &response.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

func (u *CategoryUsecase) UpdateCategory(id uint, name string) (*domain.Category, error) {
	category, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	category.Name = name
	if err := u.repo.Update(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (u *CategoryUsecase) DeleteCategory(id uint) error {
	return u.repo.Delete(id)
}
