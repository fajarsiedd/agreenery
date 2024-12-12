package request

import "go-agreenery/entities"

type NotificationRequest struct {
	ID        string
	Title     string `json:"title" validate:"required"`
	Subtitle  string `json:"subtitle" validate:"required"`
	SendAt    string `json:"send_at" validate:"required"`
	ActionURL string `json:"action_url"`
}

func (r NotificationRequest) ToEntity() entities.Notification {
	return entities.Notification{
		Base:      entities.Base{ID: r.ID},
		Title:     r.Title,
		Subtitle:  r.Subtitle,
		ActionURL: r.ActionURL,
	}
}
