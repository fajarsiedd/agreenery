package entities

type PostReport struct {
	Base
	UserID     string
	User       User
	PostID     string
	ReportType string
	StatusDone bool
}
