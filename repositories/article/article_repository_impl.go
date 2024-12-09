package article

import (
	"go-agreenery/entities"
	"go-agreenery/models"

	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{
		db: db,
	}
}

func (r articleRepository) GetArticles(filter entities.Filter) ([]entities.Article, entities.Pagination, error) {
	articleModel := models.ListArticle{}

	query := r.db.Model(&articleModel)

	if filter.PublishStatus != "" {
		if filter.PublishStatus == "false" {
			query = query.Where("publish_status = ?", false)
		}
		if filter.PublishStatus == "true" {
			query = query.Where("publish_status = ?", true)
		}
	}

	if filter.Category != "" {
		query = query.InnerJoins("Category").Where("Category.Name = ?", filter.Category)
	}

	if filter.Search != "" {
		query = query.Table("articles").Where("articles.title LIKE ?", "%"+filter.Search+"%")
	}

	if !filter.StartDate.IsZero() && !filter.EndDate.IsZero() {
		query = query.Where("articles.created_at BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	}

	query = query.Order(filter.SortBy + " " + filter.Sort)

	var totalItems int64

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	offset := (filter.Page - 1) * filter.Limit

	if err := query.Preload("Category").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Limit(filter.Limit).Offset(offset).Find(&articleModel).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	pagination := entities.Pagination{
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalItems: int(totalItems),
		TotalPages: int((int(totalItems) + filter.Limit - 1) / filter.Limit),
	}

	return articleModel.ToListEntity(), pagination, nil
}

func (r articleRepository) GetArticle(id string) (entities.Article, error) {
	articleModel := models.Article{}

	if err := r.db.Preload("Category").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).First(&articleModel, &id).Error; err != nil {
		return entities.Article{}, err
	}

	return articleModel.ToEntity(), nil
}

func (r articleRepository) CreateArticle(article entities.Article) (entities.Article, error) {
	articleModel := models.Article{}.FromEntity(article)

	if err := r.db.Create(&articleModel).Preload("Category").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Find(&articleModel).Error; err != nil {
		return entities.Article{}, err
	}

	return articleModel.ToEntity(), nil
}

func (r articleRepository) UpdateArticle(article entities.Article) (entities.Article, error) {
	articleModel := models.Article{}.FromEntity(article)

	if err := r.db.Updates(&articleModel).Preload("Category").Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Find(&articleModel).Error; err != nil {
		return entities.Article{}, err
	}

	return articleModel.ToEntity(), nil
}

func (r articleRepository) DeleteArticle(id string) (string, error) {
	articleModel := models.Article{}

	var media string
	err := r.db.Transaction(func(tx *gorm.DB) error {
		articleDb := models.Article{}
		if err := tx.First(&articleDb, &id).Error; err != nil {
			return err
		}

		media = articleDb.Thumbnail

		if err := tx.Unscoped().Delete(&articleModel, &id).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return media, nil
}
