package entities

import "time"

type Notification struct {
	Base
	UserID    string
	User      User
	Title     string
	Subtitle  string
	ActionURL string
	SendAt    time.Time
	IsSent    bool
}
