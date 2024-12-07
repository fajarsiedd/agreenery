package step

import (
	"go-agreenery/entities"
	"go-agreenery/models"

	"gorm.io/gorm"
)

type stepRepository struct {
	db *gorm.DB
}

func NewStepRepository(db *gorm.DB) *stepRepository {
	return &stepRepository{
		db: db,
	}
}

func (r stepRepository) CreateStep(step entities.Step) (entities.Plant, error) {
	stepModel := models.Step{}.FromEntity(step)
	plantModel := models.Plant{}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&stepModel).Error; err != nil {
			return err
		}

		if err := tx.Preload("Category").Preload("Steps", func(db *gorm.DB) *gorm.DB {
			db = db.Order("created_at ASC")
			return db
		}).First(&plantModel, &stepModel.PlantID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entities.Plant{}, err
	}

	return plantModel.ToEntity(), nil
}

func (r stepRepository) UpdateStep(step entities.Step) (entities.Plant, error) {
	stepModel := models.Step{}.FromEntity(step)
	plantModel := models.Plant{}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Updates(&stepModel).Error; err != nil {
			return err
		}

		if err := tx.Preload("Category").Preload("Steps", func(db *gorm.DB) *gorm.DB {
			db = db.Order("created_at ASC")
			return db
		}).First(&plantModel, &step.PlantID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entities.Plant{}, err
	}

	return plantModel.ToEntity(), nil
}

func (r stepRepository) DeleteStep(id string) (entities.Plant, error) {
	stepModel := models.Step{}
	plantModel := models.Plant{}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&stepModel, &id).Error; err != nil {
			return err
		}

		if err := tx.Unscoped().Delete(&stepModel, &id).Error; err != nil {
			return err
		}

		if err := tx.Preload("Category").Preload("Steps", func(db *gorm.DB) *gorm.DB {
			db = db.Order("created_at ASC")
			return db
		}).First(&plantModel, &stepModel.PlantID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entities.Plant{}, err
	}

	return plantModel.ToEntity(), nil
}
