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

func (u *CategoryUsecase) GetCategories() ([]response.CategoryResponse, error) {
	categories, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var categoryResponses []response.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, response.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return categoryResponses, nil
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

func (u *CategoryUsecase) UpdateCategory(id uint, name string) error {
	category, err := u.repo.FindByID(id)
	if err != nil {
		return err
	}
	category.Name = name
	if err := u.repo.Update(category); err != nil {
		return err
	}
	return nil
}

func (u *CategoryUsecase) DeleteCategory(id uint) error {
	return u.repo.Delete(id)
}
