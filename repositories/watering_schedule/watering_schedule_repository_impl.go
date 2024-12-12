package watering_schedule

import (
	"go-agreenery/constants"
	"go-agreenery/entities"
	"go-agreenery/models"

	"gorm.io/gorm"
)

type wateringScheduleRepository struct {
	db *gorm.DB
}

func NewWateringScheduleRepository(db *gorm.DB) *wateringScheduleRepository {
	return &wateringScheduleRepository{
		db: db,
	}
}

func (r wateringScheduleRepository) GetWateringSchedules(filter entities.Filter) ([]entities.WateringSchedule, entities.Pagination, error) {
	wateringScheduleModel := models.ListWateringSchedule{}

	query := r.db.Debug().Model(&wateringScheduleModel)

	if filter.Search != "" {
		query = query.Table("watering_schedules").Where("watering_schedules.plant_name LIKE ?", "%"+filter.Search+"%")
	}

	if !filter.StartDate.IsZero() && !filter.EndDate.IsZero() {
		query = query.Where("watering_schedules.created_at BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	}

	query = query.Order("watering_schedules." + filter.SortBy + " " + filter.Sort)

	var totalItems int64

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	offset := (filter.Page - 1) * filter.Limit

	if err := query.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Limit(filter.Limit).Offset(offset).Find(&wateringScheduleModel).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	pagination := entities.Pagination{
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalItems: int(totalItems),
		TotalPages: int((int(totalItems) + filter.Limit - 1) / filter.Limit),
	}

	return wateringScheduleModel.ToListEntity(), pagination, nil
}

func (r wateringScheduleRepository) GetWateringSchedule(id string) (entities.WateringSchedule, error) {
	wateringScheduleModel := models.WateringSchedule{}

	if err := r.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).First(&wateringScheduleModel, &id).Error; err != nil {
		return entities.WateringSchedule{}, err
	}

	return wateringScheduleModel.ToEntity(), nil
}

func (r wateringScheduleRepository) CreateWateringSchedule(schedule entities.WateringSchedule) (entities.WateringSchedule, error) {
	wateringScheduleModel := models.WateringSchedule{}.FromEntity(schedule)

	if err := r.db.Create(&wateringScheduleModel).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Find(&wateringScheduleModel).Error; err != nil {
		return entities.WateringSchedule{}, err
	}

	return wateringScheduleModel.ToEntity(), nil
}

func (r wateringScheduleRepository) UpdateWateringSchedule(schedule entities.WateringSchedule, currUserID string) (entities.WateringSchedule, error) {
	wateringScheduleModel := models.WateringSchedule{}.FromEntity(schedule)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		wateringScheduleDb := models.WateringSchedule{}
		if err := tx.First(&wateringScheduleDb, &schedule.ID).Error; err != nil {
			return err
		}

		if wateringScheduleDb.UserID != currUserID {
			return constants.ErrAccessNotAllowed
		}

		if err := tx.Updates(&wateringScheduleModel).Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Credential")
		}).Find(&wateringScheduleModel).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return entities.WateringSchedule{}, err
	}

	return wateringScheduleModel.ToEntity(), nil
}

func (r wateringScheduleRepository) DeleteWateringSchedule(id, currUserID string) (string, error) {
	var image string
	err := r.db.Transaction(func(tx *gorm.DB) error {
		wateringScheduleDb := models.WateringSchedule{}
		if err := tx.First(&wateringScheduleDb, &id).Error; err != nil {
			return err
		}

		if wateringScheduleDb.UserID != currUserID {
			return constants.ErrAccessNotAllowed
		}

		image = wateringScheduleDb.Image

		if err := tx.Unscoped().Delete(&wateringScheduleDb).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return "", err
	}

	return image, nil
}
