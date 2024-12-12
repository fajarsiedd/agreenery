package notification

import "github.com/labstack/echo/v4"

type NotificationHandler interface {
	GetNotifications(c echo.Context) error
	CreateNotification(c echo.Context) error
	UpdateNotification(c echo.Context) error
	DeleteNotification(c echo.Context) error
	SendNotification(c echo.Context) error
}
