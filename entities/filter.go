package entities

import "time"

type Filter struct {
	Page          int
	Limit         int
	Search        string
	Sort          string
	SortBy        string
	StartDate     time.Time
	EndDate       time.Time
	UserID        string
	Category      string
	CategoryType  string
	PostID        string
	PublishStatus string	
}
