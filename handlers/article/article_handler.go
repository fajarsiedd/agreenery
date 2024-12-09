package article

import "github.com/labstack/echo/v4"

type ArticleHandler interface {
	GetArticles(c echo.Context) error
	GetArticle(c echo.Context) error
	CreateArticle(c echo.Context) error
	UpdateArticle(c echo.Context) error
	DeleteArticle(c echo.Context) error
}
