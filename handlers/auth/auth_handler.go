package auth

import "github.com/labstack/echo/v4"

type AuthHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	GetNewTokens(c echo.Context) error
	GetProfile(c echo.Context) error
	UpdateProfile(c echo.Context) error
	UploadProfilePhoto(c echo.Context) error
}
