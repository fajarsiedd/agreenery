package models

import (
	"go-agreenery/entities"
)

type Post struct {
	Base
	Content       string
	Media         string
	UserID        string   `gorm:"size:191"`
	User          User     `gorm:"foreignKey:UserID;references:ID"`
	CategoryID    string   `gorm:"size:191"`
	Category      Category `gorm:"foreignKey:CategoryID;references:ID"`
	Comments      ListComment
	CountComments int64 `gorm:"->;-:migration"`
	Likes         ListLike
	CountLikes    int64 `gorm:"->;-:migration"`
	IsLiked       bool  `gorm:"->;-:migration"`
}

type ListPost []Post

func (p Post) FromEntity(post entities.Post) Post {
	return Post{
		Base:       p.Base.FromEntity(post.Base),
		Content:    post.Content,
		Media:      post.Media,
		UserID:     post.UserID,
		User:       p.User.FromEntity(post.User),
		CategoryID: post.CategoryID,
		Category:   p.Category.FromEntity(post.Category),
		Comments:   p.Comments.FromListEntity(post.Comments),
		Likes:      p.Likes.FromListEntity(post.Likes),
	}
}

func (p Post) ToEntity() entities.Post {
	return entities.Post{
		Base:          p.Base.ToEntity(),
		Content:       p.Content,
		Media:         p.Media,
		UserID:        p.UserID,
		User:          p.User.ToEntity(),
		CategoryID:    p.CategoryID,
		Category:      p.Category.ToEntity(),
		Comments:      p.Comments.ToListEntity(),
		CountComments: p.CountComments,
		Likes:         p.Likes.ToListEntity(),
		CountLikes:    p.CountLikes,
		IsLiked:       p.IsLiked,
	}
}

func (lp ListPost) FromListEntity(categories []entities.Post) ListPost {
	data := ListPost{}

	for _, v := range categories {
		data = append(data, Post{}.FromEntity(v))
	}

	return data
}

func (lp ListPost) ToListEntity() []entities.Post {
	data := []entities.Post{}

	for _, v := range lp {
		data = append(data, v.ToEntity())
	}

	return data
}
