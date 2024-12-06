package handler

import (
	"github.com/labstack/echo/v4"
)

type CategoryHandler interface {
	GetCategories(c echo.Context) error
	GetCategory(c echo.Context) error
	CreateCategory(c echo.Context) error
	UpdateCategory(c echo.Context) error
	DeleteCategory(c echo.Context) error
}
