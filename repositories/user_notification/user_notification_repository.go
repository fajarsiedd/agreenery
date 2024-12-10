package user_notification

import "go-agreenery/entities"

type UserNotificationRepository interface {
	GetUserNotifications(filter entities.Filter) ([]entities.UserNotification, entities.Pagination, error)
	CreateUserNotification(userNotification entities.UserNotification) (entities.UserNotification, error)
	DeleteNotification(id, currUserID string) error
	MarkNotificationAsRead(id, currUserID string) error
	MarkAllNotificationsAsRead(currUserID string) error
}
