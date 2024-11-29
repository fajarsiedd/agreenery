package auth

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error	
	GetNewTokens(c echo.Context) error
}
