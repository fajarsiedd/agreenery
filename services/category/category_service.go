package category

import "go-agreenery/entities"

type CategoryService interface {
	GetCategories(filter entities.Filter) ([]entities.Category, entities.Pagination, error)
	GetCategory(id string) (entities.Category, error)
	CreateCategory(category entities.Category) (entities.Category, error)
	UpdateCategory(category entities.Category) (entities.Category, error)
	DeleteCategory(id string) error
}