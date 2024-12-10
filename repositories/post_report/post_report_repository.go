package post_report

import "go-agreenery/entities"

type PostReportRepository interface {
	GetPostReports(filter entities.Filter) ([]entities.PostReport, entities.Pagination, error)
	CreatePostReport(postReport entities.PostReport) (entities.PostReport, error)
	DeletePostReport(id string) error
}
