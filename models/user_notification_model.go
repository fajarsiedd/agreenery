package models

import (
	"database/sql"
	"go-agreenery/entities"
)

type UserNotification struct {
	Base
	UserID    string         `gorm:"size:191"`
	Title     string         `gorm:"size:255"`
	Subtitle  string         `gorm:"size:255"`
	ActionURL string         `gorm:"size:255"`
	IsRead    bool           `gorm:"default:false"`
	PostID    sql.NullString `gorm:"size:191"`
	CommentID sql.NullString `gorm:"size:191"`
	LikeID    sql.NullString `gorm:"size:191"`
	Icon      string         `gorm:"size:255"`
}

type ListUserNotification []UserNotification

func (n UserNotification) FromEntity(userNotification entities.UserNotification) UserNotification {
	return UserNotification{
		Base:      n.Base.FromEntity(userNotification.Base),
		UserID:    userNotification.UserID,
		Title:     userNotification.Title,
		Subtitle:  userNotification.Subtitle,
		ActionURL: userNotification.ActionURL,
		IsRead:    userNotification.IsRead,
		PostID:    userNotification.PostID,
		LikeID:    userNotification.LikeID,
		CommentID: userNotification.CommentID,
		Icon:      userNotification.Icon,
	}
}

func (n UserNotification) ToEntity() entities.UserNotification {
	return entities.UserNotification{
		Base:      n.Base.ToEntity(),
		UserID:    n.UserID,
		Title:     n.Title,
		Subtitle:  n.Subtitle,
		ActionURL: n.ActionURL,
		IsRead:    n.IsRead,
		PostID:    n.PostID,
		LikeID:    n.LikeID,
		CommentID: n.CommentID,
		Icon:      n.Icon,
	}
}

func (ln ListUserNotification) FromListEntity(userNotifications []entities.UserNotification) ListUserNotification {
	data := ListUserNotification{}

	for _, v := range userNotifications {
		data = append(data, UserNotification{}.FromEntity(v))
	}

	return data
}

func (ln ListUserNotification) ToListEntity() []entities.UserNotification {
	data := []entities.UserNotification{}

	for _, v := range ln {
		data = append(data, v.ToEntity())
	}

	return data
}
