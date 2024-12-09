package comment

import (
	"go-agreenery/constants"
	"go-agreenery/entities"
	"go-agreenery/models"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r commentRepository) GetComments(filter entities.Filter) ([]entities.Comment, entities.Pagination, error) {
	commentModel := models.ListComment{}

	if err := r.db.First(&models.Post{}, &filter.PostID).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	query := r.db.Model(&commentModel).Where("post_id = ?", filter.PostID)

	if filter.Search != "" {
		query = query.Where("message LIKE ?", "%"+filter.Search+"%")
	}

	if !filter.StartDate.IsZero() && !filter.EndDate.IsZero() {
		query = query.Where("comments.created_at BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	}

	query = query.Order("comments." + filter.SortBy + " " + filter.Sort)

	var totalItems int64
	if err := query.Count(&totalItems).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	offset := (filter.Page - 1) * filter.Limit

	query = query.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	})

	if err := query.Limit(filter.Limit).Offset(offset).Find(&commentModel).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	pagination := entities.Pagination{
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalItems: int(totalItems),
		TotalPages: int((int(totalItems) + filter.Limit - 1) / filter.Limit),
	}

	return commentModel.ToListEntity(), pagination, nil
}

func (r commentRepository) CreateComment(comment entities.Comment) (entities.Comment, error) {
	commentModel := models.Comment{}.FromEntity(comment)

	if err := r.db.First(&models.Post{}, &comment.PostID).Error; err != nil {
		return entities.Comment{}, err
	}

	if err := r.db.Create(&commentModel).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Find(&commentModel).Error; err != nil {
		return entities.Comment{}, err
	}

	return commentModel.ToEntity(), nil
}

func (r commentRepository) UpdateComment(comment entities.Comment, currUserID string) (entities.Comment, error) {
	commentModel := models.Comment{}.FromEntity(comment)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&models.Post{}, &comment.PostID).Error; err != nil {
			return err
		}

		commentDb := models.Comment{}
		if err := tx.First(&commentDb, &comment.ID).Error; err != nil {
			return err
		}

		if commentDb.UserID != currUserID {
			return constants.ErrAccessNotAllowed
		}

		if err := tx.Updates(&commentModel).Error; err != nil {
			return err
		}

		if err := tx.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Credential")
		}).Find(&commentModel).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return entities.Comment{}, err
	}

	return commentModel.ToEntity(), nil
}

func (r commentRepository) DeleteComment(id, postID, currUserID string) error {
	commentModel := models.Comment{}

	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&models.Post{}, &postID).Error; err != nil {
			return err
		}

		commentDb := models.Comment{}
		if err := tx.First(&commentDb, &id).Error; err != nil {
			return err
		}

		if commentDb.UserID != currUserID {
			return constants.ErrAccessNotAllowed
		}

		if err := tx.Unscoped().Delete(&commentModel, &id).Error; err != nil {
			return err
		}

		return nil
	})
}
