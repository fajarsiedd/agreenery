package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
	"time"
)

type NotificationResponse struct {
	base.Base
	Title     string          `json:"title"`
	Subtitle  string          `json:"subtitle"`
	User      ProfileResponse `json:"user"`
	SendAt    time.Time       `json:"send_at"`
	IsSent    bool            `json:"is_sent"`
	ActionURL string          `json:"action_url"`
}

type ListNotificationResponse []NotificationResponse

func (r NotificationResponse) FromEntity(notification entities.Notification) NotificationResponse {
	return NotificationResponse{
		Base:      r.Base.FromEntity(notification.Base),
		User:      r.User.FromEntity(notification.User),
		Title:     notification.Title,
		Subtitle:  notification.Subtitle,
		ActionURL: notification.ActionURL,
		SendAt:    notification.SendAt,
		IsSent:    notification.IsSent,
	}
}

func (lr ListNotificationResponse) FromListEntity(notifications []entities.Notification) ListNotificationResponse {
	data := ListNotificationResponse{}

	for _, v := range notifications {
		data = append(data, NotificationResponse{}.FromEntity(v))
	}

	return data
}
