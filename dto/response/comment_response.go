package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type CommentResponse struct {
	base.Base
	User    ProfileResponse `json:"user"`
	PostID  string          `json:"post_id"`
	Message string          `json:"message"`
}

type ListCommentResponse []CommentResponse

func (r CommentResponse) FromEntity(comment entities.Comment) CommentResponse {
	return CommentResponse{
		Base:    r.Base.FromEntity(comment.Base),
		User:    r.User.FromEntity(comment.User),
		PostID:  comment.PostID,
		Message: comment.Message,
	}
}

func (lr ListCommentResponse) FromListEntity(comments []entities.Comment) ListCommentResponse {
	data := ListCommentResponse{}

	for _, v := range comments {
		data = append(data, CommentResponse{}.FromEntity(v))
	}

	return data
}
