package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type PostReportResponse struct {
	base.Base
	User       ProfileResponse `json:"user"`
	PostID     string          `json:"post_id"`
	ReportType string          `json:"report_type"`
	StatusDone bool            `json:"status_done"`
}

type ListPostReportResponse []PostReportResponse

func (r PostReportResponse) FromEntity(postReport entities.PostReport) PostReportResponse {
	return PostReportResponse{
		Base:       r.Base.FromEntity(postReport.Base),
		User:       r.User.FromEntity(postReport.User),
		PostID:     postReport.PostID,
		ReportType: postReport.ReportType,
		StatusDone: postReport.StatusDone,
	}
}

func (lr ListPostReportResponse) FromListEntity(postReports []entities.PostReport) ListPostReportResponse {
	data := ListPostReportResponse{}

	for _, v := range postReports {
		data = append(data, PostReportResponse{}.FromEntity(v))
	}

	return data
}
