package post

import "go-agreenery/entities"

type PostRepository interface {
	GetPosts(filter entities.Filter) ([]entities.Post, entities.Pagination, error)
	GetPost(id, userID string) (entities.Post, error)
	CreatePost(category entities.Post) (entities.Post, error)
	UpdatePost(category entities.Post, currUserID string) (entities.Post, error)
	DeletePost(id, currUserID string) (string, error)
}
