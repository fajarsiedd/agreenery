package request

import "go-agreenery/entities"

type LikeRequest struct {
	UserID string
	PostID string `json:"post_id"`
}

func (r LikeRequest) ToEntitiy() entities.Like {
	return entities.Like{
		UserID: r.UserID,
		PostID: r.PostID,
	}
}
