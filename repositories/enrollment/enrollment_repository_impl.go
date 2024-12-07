package enrollment

import (
	"go-agreenery/entities"
	"go-agreenery/models"

	"gorm.io/gorm"
)

type enrollmentRepository struct {
	db *gorm.DB
}

func NewEnrollmentRepository(db *gorm.DB) *enrollmentRepository {
	return &enrollmentRepository{
		db: db,
	}
}

func (r enrollmentRepository) CreateEnrollment(plant entities.EnrolledPlant) (entities.EnrolledPlant, error) {
	enrolledPlantModel := models.EnrolledPlant{}.FromEntity(plant)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		plantModel := models.Plant{}
		if err := tx.Preload("Steps", func(db *gorm.DB) *gorm.DB {
			db = db.Order("created_at ASC")
			return db
		}).First(&plantModel, &plant.PlantID).Error; err != nil {
			return err
		}

		if err := tx.Omit("Plant").Create(&enrolledPlantModel).Error; err != nil {
			return err
		}

		var enrolledStepModel models.EnrolledStep
		for _, v := range plantModel.Steps {
			enrolledStepModel = models.EnrolledStep{
				UserID:          enrolledPlantModel.UserID,
				EnrolledPlantID: enrolledPlantModel.ID,
				StepID:          v.ID,
			}

			if err := tx.Omit("Step").Create(&enrolledStepModel).Error; err != nil {
				return err
			}
		}

		if err := tx.Preload("Plant", func(db *gorm.DB) *gorm.DB {
			db = db.Preload("Category")
			return db
		}).Preload("EnrolledSteps", func(db *gorm.DB) *gorm.DB {
			db = db.Preload("Step").Order("created_at ASC")
			return db
		}).Find(&enrolledPlantModel).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entities.EnrolledPlant{}, err
	}

	return enrolledPlantModel.ToEntity(), nil
}

func (r enrollmentRepository) GetEnrollments(filter entities.Filter) ([]entities.EnrolledPlant, entities.Pagination, error) {
	enrolledPlantModel := models.ListEnrolledPlant{}

	query := r.db.Model(&enrolledPlantModel).Where("user_id = ?", filter.UserID)

	if filter.Search != "" {
		query = query.InnerJoins("Plant").Where("Plant.Name LIKE ?", "%"+filter.Search+"%")
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

	if err := query.Preload("Plant", func(db *gorm.DB) *gorm.DB {
		db = db.Preload("Category")

		return db
	}).Limit(filter.Limit).Offset(offset).Find(&enrolledPlantModel).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	pagination := entities.Pagination{
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalItems: int(totalItems),
		TotalPages: int((int(totalItems) + filter.Limit - 1) / filter.Limit),
	}

	return enrolledPlantModel.ToListEntity(), pagination, nil
}

func (r enrollmentRepository) GetEnrollment(enrollmentID string) (entities.EnrolledPlant, error) {
	enrolledPlantModel := models.EnrolledPlant{}

	if err := r.db.Preload("Plant", func(db *gorm.DB) *gorm.DB {
		db = db.Preload("Category")
		return db
	}).Preload("EnrolledSteps", func(db *gorm.DB) *gorm.DB {
		db = db.Preload("Step").Order("created_at ASC")
		return db
	}).Where("id = ?", enrollmentID).Find(&enrolledPlantModel).Error; err != nil {
		return entities.EnrolledPlant{}, err
	}

	return enrolledPlantModel.ToEntity(), nil
}

func (r enrollmentRepository) MarkStepAsComplete(stepID string) (entities.EnrolledPlant, error) {
	enrolledPlantModel := models.EnrolledPlant{}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		enrolledStepModel := models.EnrolledStep{}

		if err := tx.Model(&enrolledStepModel).Where("id = ?", stepID).Update("mark_complete", true).First(&enrolledStepModel, &stepID).Error; err != nil {
			return err
		}

		if err := tx.Preload("Plant", func(db *gorm.DB) *gorm.DB {
			db = db.Preload("Category")
			return db
		}).Preload("EnrolledSteps", func(db *gorm.DB) *gorm.DB {
			db = db.Preload("Step").Order("created_at ASC")
			return db
		}).First(&enrolledPlantModel, &enrolledStepModel.EnrolledPlantID).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return entities.EnrolledPlant{}, err
	}

	return enrolledPlantModel.ToEntity(), nil
}

func (r enrollmentRepository) SetEnrollmentStatusAsDone(enrollmentID string) (entities.EnrolledPlant, error) {
	enrolledPlantModel := models.EnrolledPlant{}

	if err := r.db.Where("id = ?", enrollmentID).Model(&enrolledPlantModel).Update("is_done", true).Preload("Plant", func(db *gorm.DB) *gorm.DB {
		db = db.Preload("Category")
		return db
	}).Preload("EnrolledSteps", func(db *gorm.DB) *gorm.DB {
		db = db.Preload("Step").Order("created_at ASC")
		return db
	}).Find(&enrolledPlantModel).Error; err != nil {
		return entities.EnrolledPlant{}, err
	}

	return enrolledPlantModel.ToEntity(), nil
}

func (r enrollmentRepository) DeleteEnrollment(enrollmentID string) error {
	enrolledPlantModel := models.EnrolledPlant{}
	enrolledStepModel := models.EnrolledStep{}

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Unscoped().Where("enrolled_plant_id = ?", enrollmentID).Delete(&enrolledStepModel).Error; err != nil {
			return err
		}

		if err := r.db.Unscoped().Delete(&enrolledPlantModel, &enrollmentID).Error; err != nil {
			return err
		}

		return nil
	})
}
