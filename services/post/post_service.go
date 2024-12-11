package post

import "go-agreenery/entities"

type PostService interface {
	GetPosts(filter entities.Filter) ([]entities.Post, entities.Pagination, error)
	GetPost(id, userID string) (entities.Post, error)
	CreatePost(post entities.Post) (entities.Post, error)
	UpdatePost(post entities.Post, currUserID string) (entities.Post, error)
	DeletePost(id, currUserID string, isAdmin bool) error
	LikePost(id, currUserID string) (entities.Post, error)
	GetPostsCountByCategory() ([]entities.Category, error)
}
