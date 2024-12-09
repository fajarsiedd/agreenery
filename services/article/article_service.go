package article

import "go-agreenery/entities"

type ArticleService interface {
	GetArticles(filter entities.Filter) ([]entities.Article, entities.Pagination, error)
	GetArticle(id string) (entities.Article, error)
	CreateArticle(article entities.Article) (entities.Article, error)
	UpdateArticle(article entities.Article) (entities.Article, error)
	DeleteArticle(id string) error
}
