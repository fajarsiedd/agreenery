package request

import "go-agreenery/entities"

type PostReportRequest struct {
	PostID     string `json:"post_id"`
	ReportType string `json:"report_type" validate:"required"`
}

func (r PostReportRequest) ToEntity() entities.PostReport {
	return entities.PostReport{
		PostID:     r.PostID,
		ReportType: r.ReportType,
	}
}
