package user_notification

import "github.com/labstack/echo/v4"

type UserNotificationHandler interface {
	GetUserNotifications(c echo.Context) error
	DeleteNotification(c echo.Context) error
	MarkNotificationAsRead(c echo.Context) error
	MarkAllNotificationsAsRead(c echo.Context) error
}
