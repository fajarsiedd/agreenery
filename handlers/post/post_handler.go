package post

import "github.com/labstack/echo/v4"

type PostHandler interface {
	GetPosts(c echo.Context) error
	GetPost(c echo.Context) error
	CreatePost(c echo.Context) error
	UpdatePost(c echo.Context) error
	DeletePost(c echo.Context) error
	LikePost(c echo.Context) error
}
