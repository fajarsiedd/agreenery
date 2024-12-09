package request

import (
	"go-agreenery/entities"
	"mime/multipart"
)

type ArticleRequest struct {
	ID            string
	Title         string         `form:"title" validate:"required"`
	Content       string         `form:"content" validate:"required"`
	Thumbnail     multipart.File `form:"media"`
	CategoryID    string         `form:"category_id" validate:"required"`
	PublishStatus bool           `form:"publish_status"`
}

func (r ArticleRequest) ToEntity() entities.Article {
	return entities.Article{
		Base: entities.Base{
			ID: r.ID,
		},
		Title:         r.Title,
		Content:       r.Content,
		ThumbnailFile: r.Thumbnail,
		CategoryID:    r.CategoryID,
		PublishStatus: r.PublishStatus,
	}
}
