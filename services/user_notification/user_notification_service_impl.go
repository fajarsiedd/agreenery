package user_notification

import (
	"go-agreenery/entities"
	"go-agreenery/repositories/user_notification"
)

type userNotificationService struct {
	repository user_notification.UserNotificationRepository
}

func NewUserNotificationService(r user_notification.UserNotificationRepository) *userNotificationService {
	return &userNotificationService{
		repository: r,
	}
}

func (s userNotificationService) GetUserNotifications(filter entities.Filter) ([]entities.UserNotification, entities.Pagination, error) {
	return s.repository.GetUserNotifications(filter)
}

func (s userNotificationService) CreateUserNotification(userNotification entities.UserNotification) (entities.UserNotification, error) {
	return s.repository.CreateUserNotification(userNotification)
}

func (s userNotificationService) DeleteNotification(id, currUserID string) error {
	return s.repository.DeleteNotification(id, currUserID)
}

func (s userNotificationService) MarkNotificationAsRead(id, currUserID string) error {
	return s.repository.MarkNotificationAsRead(id, currUserID)
}

func (s userNotificationService) MarkAllNotificationsAsRead(currUserID string) error {
	return s.repository.MarkAllNotificationsAsRead(currUserID)
}
