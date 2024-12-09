package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type ArticleResponse struct {
	base.Base
	Title         string          `json:"title"`
	Content       string          `json:"content"`
	Category      string          `json:"category"`
	User          ProfileResponse `json:"user"`
	Thumbnail     string          `json:"thumbnail"`
	PublishStatus bool            `json:"publish_status"`
}

type ListArticleResponse []ArticleResponse

func (r ArticleResponse) FromEntity(article entities.Article) ArticleResponse {
	return ArticleResponse{
		Base:          r.Base.FromEntity(article.Base),
		User:          r.User.FromEntity(article.User),
		Category:      article.Category.Name,
		Title:         article.Title,
		Content:       article.Content,
		Thumbnail:     article.Thumbnail,
		PublishStatus: article.PublishStatus,
	}
}

func (lr ListArticleResponse) FromListEntity(articles []entities.Article) ListArticleResponse {
	data := ListArticleResponse{}

	for _, v := range articles {
		data = append(data, ArticleResponse{}.FromEntity(v))
	}

	return data
}
