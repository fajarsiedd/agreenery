package post_report

import (
	"go-agreenery/entities"
	"go-agreenery/helpers"
	"go-agreenery/repositories/post_report"
	"strings"
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

func (s postReportService) DeletePostWithMessage(postReport entities.PostReport) error {
	media, err := s.repository.DeletePostWithMessage(postReport)
	if err != nil {
		return err
	}

	if media != "" {
		var oldObj string
		splittedStr := strings.Split(media, "/")
		oldObj = splittedStr[len(splittedStr)-1]

		if err := helpers.DeleteFile(oldObj); err != nil {
			return err
		}
	}

	return nil
}

func (s postReportService) SendWarning(postReport entities.PostReport) error {
	return s.repository.SendWarning(postReport)
}
