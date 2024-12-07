package plant

import (
	"go-agreenery/entities"
	"go-agreenery/models"

	"gorm.io/gorm"
)

type plantRepository struct {
	db *gorm.DB
}

func NewPlantRepository(db *gorm.DB) *plantRepository {
	return &plantRepository{
		db: db,
	}
}

func (r plantRepository) GetPlants(filter entities.Filter) ([]entities.Plant, entities.Pagination, error) {
	plantModel := models.ListPlant{}

	query := r.db.Model(&plantModel)

	if filter.Category != "" {
		query = query.Joins("Category").Where("Category.Name = ?", filter.Category)
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

	if err := query.Preload("Category").Limit(filter.Limit).Offset(offset).Find(&plantModel).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	pagination := entities.Pagination{
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalItems: int(totalItems),
		TotalPages: int((int(totalItems) + filter.Limit - 1) / filter.Limit),
	}

	return plantModel.ToListEntity(), pagination, nil
}

func (r plantRepository) GetPlant(id string) (entities.Plant, error) {
	plantModel := models.Plant{}

	if err := r.db.Preload("Category").First(&plantModel, &id).Error; err != nil {
		return entities.Plant{}, err
	}

	return plantModel.ToEntity(), nil
}

func (r plantRepository) CreatePlant(plant entities.Plant) (entities.Plant, error) {
	plantModel := models.Plant{}.FromEntity(plant)

	if err := r.db.Create(&plantModel).Preload("Category").Find(&plantModel).Error; err != nil {
		return entities.Plant{}, err
	}

	return plantModel.ToEntity(), nil
}

func (r plantRepository) UpdatePlant(plant entities.Plant) (entities.Plant, error) {
	plantModel := models.Plant{}.FromEntity(plant)

	if err := r.db.Updates(&plantModel).Preload("Category").Find(&plantModel).Error; err != nil {
		return entities.Plant{}, err
	}

	return plantModel.ToEntity(), nil
}

func (r plantRepository) DeletePlant(id string) error {
	plantModel := models.Plant{}

	if err := r.db.Select("Step").Delete(&plantModel, &id).Error; err != nil {
		return err
	}

	return nil
}