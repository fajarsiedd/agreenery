package comment

import (
	"go-agreenery/entities"
	"go-agreenery/repositories/comment"
)

type commentService struct {
	repository comment.CommentRepository
}

func NewCommentService(r comment.CommentRepository) *commentService {
	return &commentService{
		repository: r,
	}
}

func (s commentService) GetComments(filter entities.Filter) ([]entities.Comment, entities.Pagination, error) {
	return s.repository.GetComments(filter)
}

func (s commentService) CreateComment(comment entities.Comment) (entities.Comment, error) {
	return s.repository.CreateComment(comment)
}

func (s commentService) UpdateComment(comment entities.Comment, currUserID string) (entities.Comment, error) {
	return s.repository.UpdateComment(comment, currUserID)
}

func (s commentService) DeleteComment(id, postID, currUserID string) error {
	return s.repository.DeleteComment(id, postID, currUserID)
}
