package notification

import (
	"go-agreenery/entities"
	"go-agreenery/models"

	"gorm.io/gorm"
)

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *notificationRepository {
	return &notificationRepository{
		db: db,
	}
}

func (r notificationRepository) GetNotifications(filter entities.Filter) ([]entities.Notification, entities.Pagination, error) {
	notificationModel := models.ListNotification{}

	query := r.db.Debug().Model(&notificationModel)

	if filter.Search != "" {
		query = query.Table("notifications").Where("notifications.title LIKE ?", "%"+filter.Search+"%")
	}

	if !filter.StartDate.IsZero() && !filter.EndDate.IsZero() {
		query = query.Where("notifications.created_at BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
	}

	query = query.Order("notifications." + filter.SortBy + " " + filter.Sort)

	var totalItems int64

	if err := query.Count(&totalItems).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	offset := (filter.Page - 1) * filter.Limit

	if err := query.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Limit(filter.Limit).Offset(offset).Find(&notificationModel).Error; err != nil {
		return nil, entities.Pagination{}, err
	}

	pagination := entities.Pagination{
		Page:       filter.Page,
		Limit:      filter.Limit,
		TotalItems: int(totalItems),
		TotalPages: int((int(totalItems) + filter.Limit - 1) / filter.Limit),
	}

	return notificationModel.ToListEntity(), pagination, nil
}

func (r notificationRepository) CreateNotification(notification entities.Notification) (entities.Notification, error) {
	notificationModel := models.Notification{}.FromEntity(notification)

	if err := r.db.Omit("User").Create(&notificationModel).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Find(&notificationModel).Error; err != nil {
		return entities.Notification{}, err
	}

	return notificationModel.ToEntity(), nil
}

func (r notificationRepository) UpdateNotification(notification entities.Notification) (entities.Notification, error) {
	notificationModel := models.Notification{}.FromEntity(notification)

	if err := r.db.Omit("User").Updates(&notificationModel).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Credential")
	}).Find(&notificationModel).Error; err != nil {
		return entities.Notification{}, err
	}

	return notificationModel.ToEntity(), nil
}

func (r notificationRepository) DeleteNotification(id string) error {
	notificationModel := models.Notification{}

	if err := r.db.Unscoped().Delete(&notificationModel, &id).Error; err != nil {
		return err
	}

	return nil
}

func (r notificationRepository) SendNotification(id string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		notificationDb := models.Notification{}
		if err := tx.First(&notificationDb, &id).Error; err != nil {
			return err
		}

		userDb := models.ListUser{}
		if err := tx.Model(&userDb).
			Joins("INNER JOIN credentials ON users.credential_id = credentials.id").Where("credentials.role = 'user'").
			Find(&userDb).Error; err != nil {
			return err
		}

		// CREATE NOITIFICATION FOR USERS
		for _, user := range userDb {
			userNotifModel := models.UserNotification{
				UserID:    user.ID,
				Title:     notificationDb.Title,
				Subtitle:  notificationDb.Subtitle,
				ActionURL: notificationDb.ActionURL,
				Icon:      "https://storage.googleapis.com/agreenery/uploads/agreenery-logo.png",
			}

			if err := tx.Create(&userNotifModel).Error; err != nil {
				return err
			}
		}

		if err := tx.Model(&notificationDb).Where("id = ?", id).Update("is_sent", true).Error; err != nil {
			return err
		}

		return nil
	})
}
