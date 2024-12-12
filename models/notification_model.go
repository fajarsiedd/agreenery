package models

import (
	"go-agreenery/entities"
	"time"
)

type Notification struct {
	Base
	UserID    string `gorm:"size:191"`
	User      User   `gorm:"foreignKey:UserID;references:ID"`
	Title     string `gorm:"size:255"`
	Subtitle  string `gorm:"size:255"`
	ActionURL string `gorm:"size:255"`
	SendAt    time.Time
	IsSent    bool `gorm:"default:false"`
}

type ListNotification []Notification

func (n Notification) FromEntity(notification entities.Notification) Notification {
	return Notification{
		Base:      n.Base.FromEntity(notification.Base),
		UserID:    notification.UserID,
		User:      n.User.FromEntity(notification.User),
		Title:     notification.Title,
		Subtitle:  notification.Subtitle,
		ActionURL: notification.ActionURL,
		SendAt:    notification.SendAt,
		IsSent:    notification.IsSent,
	}
}

func (n Notification) ToEntity() entities.Notification {
	return entities.Notification{
		Base:      n.Base.ToEntity(),
		UserID:    n.UserID,
		User:      n.User.ToEntity(),
		Title:     n.Title,
		Subtitle:  n.Subtitle,
		ActionURL: n.ActionURL,
		SendAt:    n.SendAt,
		IsSent:    n.IsSent,
	}
}

func (ln ListNotification) FromListEntity(notifications []entities.Notification) ListNotification {
	data := ListNotification{}

	for _, v := range notifications {
		data = append(data, Notification{}.FromEntity(v))
	}

	return data
}

func (ln ListNotification) ToListEntity() []entities.Notification {
	data := []entities.Notification{}

	for _, v := range ln {
		data = append(data, v.ToEntity())
	}

	return data
}
