package entities

import "database/sql"

type UserNotification struct {
	Base
	UserID    string
	Title     string
	Subtitle  string
	ActionURL string
	IsRead    bool
	PostID    sql.NullString
}
