package category

import (
	"go-agreenery/entities"
	"go-agreenery/models"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r categoryRepository) GetCategories(filter entities.Filter) ([]entities.Category, entities.Pagination, error) {
	categoryModel := models.ListCategory{}

	query := r.db.Model(&categoryModel)

	if filter.CategoryType != "" {
		query.Where("type = ?", filter.CategoryType)
	}

	if filter.Search != "" {
		query = query.Where("name LIKE ?", "%"+filter.Search+"%")
	}

	if !filter.StartDate.IsZero() && !filter.EndDate.IsZero() {
		query = query.Where("created_at BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	}

	query = query.Order(filter.SortBy + " " + filter.Sort)

	var totalItems int64

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	offset := (filter.Page - 1) * filter.Limit

	if err := query.Limit(filter.Limit).Offset(offset).Find(&categoryModel).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	pagination := entities.Pagination{
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalItems: int(totalItems),
		TotalPages: int((int(totalItems) + filter.Limit - 1) / filter.Limit),
	}

	return categoryModel.ToListEntity(), pagination, nil
}

func (r categoryRepository) GetCategory(id string) (entities.Category, error) {
	categoryModel := models.Category{}

	if err := r.db.First(&categoryModel, &id).Error; err != nil {
		return entities.Category{}, err
	}

	return categoryModel.ToEntity(), nil
}

func (r categoryRepository) CreateCategory(category entities.Category) (entities.Category, error) {
	categoryModel := models.Category{}.FromEntity(category)

	if err := r.db.Create(&categoryModel).Error; err != nil {
		return entities.Category{}, err
	}

	return categoryModel.ToEntity(), nil
}

func (r categoryRepository) UpdateCategory(category entities.Category) (entities.Category, error) {
	categoryModel := models.Category{}.FromEntity(category)

	if err := r.db.Updates(&categoryModel).Error; err != nil {
		return entities.Category{}, err
	}

	return categoryModel.ToEntity(), nil
}

func (r categoryRepository) DeleteCategory(id string) error {
	categoryModel := models.Category{}

	if err := r.db.Unscoped().Delete(&categoryModel, &id).Error; err != nil {
		return err
	}

	return nil
}
