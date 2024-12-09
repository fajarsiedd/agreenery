package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type PostResponse struct {
	base.Base
	User          ProfileResponse     `json:"user"`
	Content       string              `json:"content"`
	Category      string              `json:"category"`
	Media         string              `json:"media"`
	CountLikes    int64               `json:"count_likes"`
	CountComments int64               `json:"count_comments"`
	Comments      ListCommentResponse `json:"comments"`
	IsLiked       bool                `json:"is_liked"`
}

type ListPostResponse []PostResponse

func (r PostResponse) FromEntity(post entities.Post) PostResponse {
	return PostResponse{
		Base:          r.Base.FromEntity(post.Base),
		User:          r.User.FromEntity(post.User),
		Category:      post.Category.Name,
		Content:       post.Content,
		Media:         post.Media,
		CountLikes:    post.CountLikes,
		CountComments: post.CountComments,
		Comments:      r.Comments.FromListEntity(post.Comments),
		IsLiked:       post.IsLiked,
	}
}

func (lr ListPostResponse) FromListEntity(posts []entities.Post) ListPostResponse {
	data := ListPostResponse{}

	for _, v := range posts {
		data = append(data, PostResponse{}.FromEntity(v))
	}

	return data
}
