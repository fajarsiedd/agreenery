package post_report

import (
	"go-agreenery/entities"
	"go-agreenery/repositories/post_report"
)

type postReportService struct {
	repository post_report.PostReportRepository
}

func NewPostReportService(r post_report.PostReportRepository) *postReportService {
	return &postReportService{
		repository: r,
	}
}

func (s postReportService) GetPostReports(filter entities.Filter) ([]entities.PostReport, entities.Pagination, error) {
	return s.repository.GetPostReports(filter)
}

func (s postReportService) CreatePostReport(postReport entities.PostReport) (entities.PostReport, error) {
	return s.repository.CreatePostReport(postReport)
}

func (s postReportService) DeletePostReport(id string) error {
	return s.repository.DeletePostReport(id)
}
