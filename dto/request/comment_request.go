package request

import "go-agreenery/entities"

type CommentRequest struct {
	ID      string
	UserID  string
	Message string `json:"message" validate:"required"`
	PostID  string `json:"post_id"`
}

func (r CommentRequest) ToEntity() entities.Comment {
	return entities.Comment{
		Base: entities.Base{
			ID: r.ID,
		},
		UserID:  r.UserID,
		Message: r.Message,
		PostID:  r.PostID,
	}
}