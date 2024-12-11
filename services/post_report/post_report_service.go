package post_report

import "go-agreenery/entities"

type PostReportService interface {
	GetPostReports(filter entities.Filter) ([]entities.PostReport, entities.Pagination, error)
	CreatePostReport(postReport entities.PostReport) (entities.PostReport, error)
	DeletePostReport(id string) error
	DeletePostWithMessage(postReport entities.PostReport) error
	SendWarning(postReport entities.PostReport) error
}
