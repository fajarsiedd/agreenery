package user_notification

import (
	"go-agreenery/constants"
	"go-agreenery/entities"
	"go-agreenery/models"

	"gorm.io/gorm"
)

type userNotificationRepository struct {
	db *gorm.DB
}

func NewUserNotificationRepository(db *gorm.DB) *userNotificationRepository {
	return &userNotificationRepository{
		db: db,
	}
}

func (r userNotificationRepository) GetUserNotifications(filter entities.Filter) ([]entities.UserNotification, entities.Pagination, error) {
	userNotificationModel := models.ListUserNotification{}

	query := r.db.Model(&userNotificationModel)

	if filter.Search != "" {
		query = query.Where("user_notifications.title LIKE ?", "%"+filter.Search+"%")
	}

	if !filter.StartDate.IsZero() && !filter.EndDate.IsZero() {
		query = query.Where("user_notifications.created_at BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	}

	query = query.Order("user_notifications." + filter.SortBy + " " + filter.Sort)

	var totalItems int64

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	offset := (filter.Page - 1) * filter.Limit

	if err := query.Limit(filter.Limit).Offset(offset).Find(&userNotificationModel).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	pagination := entities.Pagination{
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalItems: int(totalItems),
		TotalPages: int((int(totalItems) + filter.Limit - 1) / filter.Limit),
	}

	return userNotificationModel.ToListEntity(), pagination, nil
}

func (r userNotificationRepository) CreateUserNotification(userNotification entities.UserNotification) (entities.UserNotification, error) {
	userNotificationModel := models.UserNotification{}.FromEntity(userNotification)

	if err := r.db.Create(&userNotificationModel).Error; err != nil {
		return entities.UserNotification{}, err
	}

	return userNotificationModel.ToEntity(), nil
}

func (r userNotificationRepository) DeleteNotification(id, currUserID string) error {
	userNotificationModel := models.UserNotification{}

	if err := r.db.First(&userNotificationModel, &id).Error; err != nil {
		return err
	}

	if userNotificationModel.UserID != currUserID {
		return constants.ErrAccessNotAllowed
	}

	if err := r.db.Unscoped().Delete(&userNotificationModel, &id).Error; err != nil {
		return err
	}

	return nil
}

func (r userNotificationRepository) MarkNotificationAsRead(id, currUserID string) error {
	userNotificationModel := models.UserNotification{}

	if err := r.db.First(&userNotificationModel, &id).Error; err != nil {
		return err
	}

	if userNotificationModel.UserID != currUserID {
		return constants.ErrAccessNotAllowed
	}

	if err := r.db.Model(&userNotificationModel).Where("id = ?", id).Update("is_read = ?", true).Error; err != nil {
		return err
	}

	return nil
}

func (r userNotificationRepository) MarkAllNotificationsAsRead(currUserID string) error {
	userNotificationModel := models.UserNotification{}

	if err := r.db.Model(&userNotificationModel).Where("user_id = ?", currUserID).Update("is_read = ?", true).Error; err != nil {
		return err
	}

	return nil
}
