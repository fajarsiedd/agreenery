package request

import (
	"go-agreenery/entities"
	"mime/multipart"
)

type PostRequest struct {
	ID         string
	UserID     string
	Content    string         `form:"content" validate:"required"`
	Media      multipart.File `form:"media"`
	CategoryID string         `form:"category_id" validate:"required"`
}

func (r PostRequest) ToEntity() entities.Post {
	return entities.Post{
		Base: entities.Base{
			ID: r.ID,
		},
		UserID:     r.UserID,
		Content:    r.Content,
		MediaFile:  r.Media,
		CategoryID: r.CategoryID,
	}
}
