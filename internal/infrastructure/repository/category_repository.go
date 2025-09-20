package repository

import (
	"kaung-htet-hein-dev/finance-tracker-go/internal/domain"
	"kaung-htet-hein-dev/finance-tracker-go/pkg"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) Create(category *domain.Category) error {
	err := r.DB.Create(category).Error

	if err != nil {
		return pkg.HandleGormError(err, "category")
	}

	return nil
}

func (r *CategoryRepository) FindAll() ([]domain.Category, error) {
	var categories []domain.Category
	result := r.DB.Find(&categories)
	return categories, result.Error
}

func (r *CategoryRepository) FindByID(id uint) (*domain.Category, error) {
	var category domain.Category
	result := r.DB.First(&category, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (r *CategoryRepository) Update(category *domain.Category) error {
	return r.DB.Save(category).Error
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.DB.Delete(&domain.Category{}, id).Error
}
