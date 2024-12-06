package category

import (
	"go-agreenery/entities"
	"go-agreenery/repositories/category"
)

type categoryService struct {
	repository category.CategoryRepository
}

func NewCategoryService(r category.CategoryRepository) *categoryService {
	return &categoryService{
		repository: r,
	}
}

func (s categoryService) GetCategories(filter entities.Filter) ([]entities.Category, entities.Pagination, error) {
	return s.repository.GetCategories(filter)
}

func (s categoryService) GetCategory(id string) (entities.Category, error) {
	return s.repository.GetCategory(id)
}

func (s categoryService) CreateCategory(category entities.Category) (entities.Category, error) {
	return s.repository.CreateCategory(category)
}

func (s categoryService) UpdateCategory(category entities.Category) (entities.Category, error) {
	return s.repository.UpdateCategory(category)
}

func (s categoryService) DeleteCategory(id string) error {
	return s.repository.DeleteCategory(id)
}
