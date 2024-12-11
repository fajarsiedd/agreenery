package post

import (
	"database/sql"
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
		query = query.Joins("INNER JOIN categories ON posts.category_id = categories.id").Where("categories.name = ?", filter.Category)
	}

	if filter.Search != "" {
		query = query.Where("posts.content LIKE ?", "%"+filter.Search+"%")
	}

	if !filter.StartDate.IsZero() && !filter.EndDate.IsZero() {
		query = query.Where("posts.created_at BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	}

	query = query.Order("posts." + filter.SortBy + " " + filter.Sort)

	var totalItems int64
	if err := query.Count(&totalItems).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	offset := (filter.Page - 1) * filter.Limit

	query = query.Preload("Category").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Preload("Comments", func(db *gorm.DB) *gorm.DB {
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

func (r postRepository) GetPost(id, currUserID string) (entities.Post, error) {
	postModel := models.Post{}

	query := r.db.Model(&postModel)

	if err := query.Preload("Category").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Select(`posts.*, 
		(SELECT COUNT(*) FROM comments WHERE comments.post_id = posts.id) AS count_comments, 
		(SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id) AS count_likes,
		EXISTS (
			SELECT 1 
			FROM likes 
			WHERE likes.post_id = posts.id AND likes.user_id = ?
		) AS is_liked
	`, currUserID).First(&postModel, &id).Error; err != nil {
		return entities.Post{}, err
	}

	return postModel.ToEntity(), nil
}

func (r postRepository) CreatePost(post entities.Post) (entities.Post, error) {
	postModel := models.Post{}.FromEntity(post)

	if err := r.db.Omit("Category", "CountComments", "CountLikes", "IsLiked").Create(&postModel).Preload("Category").Preload("User", func(db *gorm.DB) *gorm.DB {
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

		if err := tx.Omit("Category", "CountComments", "CountLikes", "IsLiked").Updates(&postModel).Error; err != nil {
			return err
		}

		if err := tx.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Credential")
		}).Preload("Category").Select(`posts.*, 
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

func (r postRepository) DeletePost(id, currUserID string, isAdmin bool) (string, error) {
	postModel := models.Post{}

	var media string
	err := r.db.Transaction(func(tx *gorm.DB) error {
		postDb := models.Post{}
		if err := tx.First(&postDb, &id).Error; err != nil {
			return err
		}

		media = postDb.Media

		if !isAdmin && postDb.UserID != currUserID {
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

func (r postRepository) LikePost(id, currUserID string) (entities.Post, error) {
	postModel := models.Post{}

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&postModel, &id).Error; err != nil {
			return err
		}

		var likeCount int64
		if err := tx.Model(&models.ListLike{}).Where("post_id = ? AND user_id = ?", id, currUserID).Count(&likeCount).Error; err != nil {
			return err
		}

		if likeCount == 0 {
			likeModel := models.Like{PostID: id, UserID: currUserID}
			if err := tx.Create(&likeModel).Error; err != nil {
				return err
			}

			if currUserID != postModel.UserID {
				// CREATE NOTIFICATION
				currUser := models.User{}
				if err := tx.First(&currUser, &currUserID).Error; err != nil {
					return err
				}

				userNotifModel := models.UserNotification{
					PostID:    sql.NullString{String: id, Valid: true},
					LikeID:    sql.NullString{String: likeModel.ID, Valid: true},
					UserID:    postModel.UserID,
					Title:     currUser.DisplayName + constants.UserLikedPost,
					Subtitle:  constants.GeneralSubtitle,
					ActionURL: id,
					Icon:      currUser.Photo,
				}
				if err := tx.Create(&userNotifModel).Error; err != nil {
					return err
				}
			}
		} else {
			likeDb := models.Like{}
			if err := tx.Where("post_id = ? AND user_id = ?", id, currUserID).Find(&likeDb).Error; err != nil {
				return err
			}

			if err := tx.Unscoped().Delete(&models.Like{}, &likeDb.ID).Error; err != nil {
				return err
			}

			if err := tx.Unscoped().Where("post_id = ? AND user_id = ? AND like_id = ?", id, postModel.UserID, likeDb.ID).Delete(&models.UserNotification{}).Error; err != nil {
				return err
			}
		}

		if err := tx.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Credential")
		}).Preload("Category").Select(`
            posts.*, 
            (SELECT COUNT(*) FROM comments WHERE comments.post_id = posts.id) AS count_comments, 
            (SELECT COUNT(*) FROM likes WHERE likes.post_id = posts.id) AS count_likes,
            EXISTS (
                SELECT 1 
                FROM likes 
                WHERE likes.post_id = posts.id AND likes.user_id = ?
            ) AS is_liked
        `, currUserID).First(&postModel, &id).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return entities.Post{}, err
	}

	return postModel.ToEntity(), nil
}

func (r postRepository) GetPostsCountByCategory() ([]entities.Category, error) {
	categoryModel := models.ListCategory{}

	query := r.db.Model(&categoryModel)

	query = query.Model(&models.Category{}).Select(`
		categories.*,
		(SELECT COUNT(*) FROM posts WHERE posts.category_id = categories.id) AS count_posts
	`).Where("type = ?", "post").Having("count_posts > 0").Order("count_posts DESC")

	if err := query.Find(&categoryModel).Error; err != nil {
		return nil, err
	}

	return categoryModel.ToListEntity(), nil
}
