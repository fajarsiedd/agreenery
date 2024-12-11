package entities

import "database/sql"

type PostReport struct {
	Base
	UserID     string
	User       User
	PostID     sql.NullString
	ReportType string
	StatusDone bool
	Message    string
}
