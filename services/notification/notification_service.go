package notification

import "go-agreenery/entities"

type NotificationService interface {
	GetNotifications(filter entities.Filter) ([]entities.Notification, entities.Pagination, error)
	CreateNotification(notification entities.Notification) (entities.Notification, error)
	UpdateNotification(notification entities.Notification) (entities.Notification, error)
	DeleteNotification(id string) error
	SendNotification(id string) error
}
