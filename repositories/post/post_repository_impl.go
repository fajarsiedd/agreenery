package post

import (
	"go-agreenery/constants"
	"go-agreenery/entities"
	"go-agreenery/models"

	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (r postRepository) GetPosts(filter entities.Filter) ([]entities.Post, entities.Pagination, error) {
	postModel := models.ListPost{}

	query := r.db.Model(&postModel)

	if filter.Category != "" {
		query = query.InnerJoins("Category").Where("Category.Name = ?", filter.Category)
	}

	if filter.Search != "" {
		query = query.Where("content LIKE ?", "%"+filter.Search+"%")
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

	query = query.Preload("Category").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Preload("Likes").Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Limit(2).Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Credential")
		})
	})

	query = query.Model(&models.Post{}).Select(`
		posts.*, 
		(SELECT COUNT(*) FROM comments WHERE comments.post_id = posts.id) AS count_comments, 
		(SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id) AS count_likes,
		EXISTS (
			SELECT 1 
			FROM likes 
			WHERE likes.post_id = posts.id AND likes.user_id = ?
		) AS is_liked
	`, filter.UserID)

	if err := query.Limit(filter.Limit).Offset(offset).Find(&postModel).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	pagination := entities.Pagination{
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalItems: int(totalItems),
		TotalPages: int((int(totalItems) + filter.Limit - 1) / filter.Limit),
	}

	return postModel.ToListEntity(), pagination, nil
}

func (r postRepository) GetPost(id, userID string) (entities.Post, error) {
	postModel := models.Post{}

	query := r.db.Model(&postModel)

	if err := query.Preload("Category").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Preload("Likes").Preload("Comments", func(db *gorm.DB) *gorm.DB {
		return db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Credential")
		})
	}).Select(`posts.*, 
		(SELECT COUNT(*) FROM comments WHERE comments.post_id = posts.id) AS count_comments, 
		(SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id) AS count_likes,
		EXISTS (
			SELECT 1 
			FROM likes 
			WHERE likes.post_id = posts.id AND likes.user_id = ?
		) AS is_liked
	`, userID).First(&postModel, &id).Error; err != nil {
		return entities.Post{}, err
	}

	return postModel.ToEntity(), nil
}

func (r postRepository) CreatePost(post entities.Post) (entities.Post, error) {
	postModel := models.Post{}.FromEntity(post)

	if err := r.db.Create(&postModel).Preload("Category").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Find(&postModel).Error; err != nil {
		return entities.Post{}, err
	}

	return postModel.ToEntity(), nil
}

func (r postRepository) UpdatePost(post entities.Post, currUserID string) (entities.Post, error) {
	postModel := models.Post{}.FromEntity(post)

	err := r.db.Transaction(func(tx *gorm.DB) error {
		postDb := models.Post{}
		if err := tx.First(&postDb, &post.ID).Error; err != nil {
			return err
		}

		if postDb.UserID != currUserID {
			return constants.ErrAccessNotAllowed
		}

		if err := tx.Updates(&postModel).Error; err != nil {
			return err
		}

		if err := tx.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Credential")
		}).Preload("Category").Preload("Likes").Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Preload("User", func(db *gorm.DB) *gorm.DB {
				return db.Preload("Credential")
			})
		}).Select(`posts.*, 
			(SELECT COUNT(*) FROM comments WHERE comments.post_id = posts.id) AS count_comments, 
			(SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id) AS count_likes,
			EXISTS (
				SELECT 1 
				FROM likes 
				WHERE likes.post_id = posts.id AND likes.user_id = ?
			) AS is_liked
		`, post.UserID).Find(&postModel).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return entities.Post{}, err
	}

	return postModel.ToEntity(), nil
}

func (r postRepository) DeletePost(id, currUserID string) (string, error) {
	postModel := models.Post{}

	var media string
	err := r.db.Transaction(func(tx *gorm.DB) error {
		postDb := models.Post{}
		if err := tx.First(&postDb, &id).Error; err != nil {
			return err
		}

		media = postDb.Media

		if postDb.UserID != currUserID {
			return constants.ErrAccessNotAllowed
		}

		if err := tx.Unscoped().Where("post_id = ?", id).Delete(&models.Comment{}).Error; err != nil {
			return err
		}

		if err := tx.Unscoped().Where("post_id = ?", id).Delete(&models.Like{}).Error; err != nil {
			return err
		}

		if err := tx.Unscoped().Delete(&postModel, &id).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return media, nil
}
