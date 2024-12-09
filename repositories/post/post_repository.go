package post

import "go-agreenery/entities"

type PostRepository interface {
	GetPosts(filter entities.Filter) ([]entities.Post, entities.Pagination, error)
	GetPost(id, currUserID string) (entities.Post, error)
	CreatePost(post entities.Post) (entities.Post, error)
	UpdatePost(post entities.Post, currUserID string) (entities.Post, error)
	DeletePost(id, currUserID string) (string, error)
}
