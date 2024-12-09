package models

import "go-agreenery/entities"

type Article struct {
	Base
	Thumbnail     string
	Title         string
	Content       string
	UserID        string   `gorm:"size:191"`
	User          User     `gorm:"foreignKey:UserID;references:ID"`
	CategoryID    string   `gorm:"size:191"`
	Category      Category `gorm:"foreignKey:CategoryID;references:ID"`
	PublishStatus bool     `gorm:"default:false"`
}

type ListArticle []Article

func (p Article) FromEntity(post entities.Article) Article {
	return Article{
		Base:          p.Base.FromEntity(post.Base),
		Title:         post.Title,
		Content:       post.Content,
		Thumbnail:     post.Thumbnail,
		PublishStatus: post.PublishStatus,
		UserID:        post.UserID,
		User:          p.User.FromEntity(post.User),
		CategoryID:    post.CategoryID,
		Category:      p.Category.FromEntity(post.Category),
	}
}

func (p Article) ToEntity() entities.Article {
	return entities.Article{
		Base:          p.Base.ToEntity(),
		Title:         p.Title,
		Content:       p.Content,
		Thumbnail:     p.Thumbnail,
		PublishStatus: p.PublishStatus,
		UserID:        p.UserID,
		User:          p.User.ToEntity(),
		CategoryID:    p.CategoryID,
		Category:      p.Category.ToEntity(),
	}
}

func (la ListArticle) FromListEntity(articles []entities.Article) ListArticle {
	data := ListArticle{}

	for _, v := range articles {
		data = append(data, Article{}.FromEntity(v))
	}

	return data
}

func (la ListArticle) ToListEntity() []entities.Article {
	data := []entities.Article{}

	for _, v := range la {
		data = append(data, v.ToEntity())
	}

	return data
}
