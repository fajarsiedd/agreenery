package request

import "go-agreenery/entities"

type PostReportActionRequest struct {
	Message string `json:"message" validate:"required"`
}

func (r PostReportActionRequest) ToEntity() entities.PostReport {
	return entities.PostReport{
		Message: r.Message,
	}
}
