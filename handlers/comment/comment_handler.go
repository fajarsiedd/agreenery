package comment

import (
	"github.com/labstack/echo/v4"
)

type CommentHandler interface {
	GetCategories(c echo.Context) error
	GetComment(c echo.Context) error
	CreateComment(c echo.Context) error
	UpdateComment(c echo.Context) error
	DeleteComment(c echo.Context) error
}
