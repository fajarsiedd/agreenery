package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type UserNotificationResponse struct {
	base.Base
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	ActionURL string `json:"action_url"`
	IsRead    bool   `json:"is_read"`
	Icon      string `json:"icon"`
}

type ListUserNotificationResponse []UserNotificationResponse

func (r UserNotificationResponse) FromEntity(userNotification entities.UserNotification) UserNotificationResponse {
	return UserNotificationResponse{
		Base:      r.Base.FromEntity(userNotification.Base),
		Title:     userNotification.Title,
		Subtitle:  userNotification.Subtitle,
		ActionURL: userNotification.ActionURL,
		IsRead:    userNotification.IsRead,
		Icon:      userNotification.Icon,
	}
}

func (lr ListUserNotificationResponse) FromListEntity(userNotifications []entities.UserNotification) ListUserNotificationResponse {
	data := ListUserNotificationResponse{}

	for _, v := range userNotifications {
		data = append(data, UserNotificationResponse{}.FromEntity(v))
	}

	return data
}
