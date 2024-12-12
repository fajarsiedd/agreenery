package notification

import (
	"go-agreenery/entities"
	"go-agreenery/repositories/notification"
)

type notificationService struct {
	repository notification.NotificationRepository
}

func NewNotificationService(r notification.NotificationRepository) *notificationService {
	return &notificationService{
		repository: r,
	}
}

func (s notificationService) GetNotifications(filter entities.Filter) ([]entities.Notification, entities.Pagination, error) {
	return s.repository.GetNotifications(filter)
}

func (s notificationService) CreateNotification(notification entities.Notification) (entities.Notification, error) {
	return s.repository.CreateNotification(notification)
}

func (s notificationService) UpdateNotification(notification entities.Notification) (entities.Notification, error) {
	return s.repository.UpdateNotification(notification)
}

func (s notificationService) DeleteNotification(id string) error {
	return s.repository.DeleteNotification(id)
}

func (s notificationService) SendNotification(id string) error {
	return s.repository.SendNotification(id)
}
