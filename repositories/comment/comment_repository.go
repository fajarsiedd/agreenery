package comment

import "go-agreenery/entities"

type CommentRepository interface {
	GetComments(filter entities.Filter) ([]entities.Comment, entities.Pagination, error)
	CreateComment(comment entities.Comment) (entities.Comment, error)
	UpdateComment(comment entities.Comment, currUserID string) (entities.Comment, error)
	DeleteComment(id, postID, currUserID string) error
}
