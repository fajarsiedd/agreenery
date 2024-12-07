package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type EnrolledStepResponse struct {
	base.Base
	Title        string `json:"name"`
	Description  string `json:"description"`
	VideoURL     string `json:"video_url"`
	MarkComplete bool   `json:"mark_complete"`
}

type ListEnrolledStepResponse []EnrolledStepResponse

func (r EnrolledStepResponse) FromEntity(step entities.EnrolledStep) EnrolledStepResponse {
	return EnrolledStepResponse{
		Base:         r.Base.FromEntity(step.Base),
		Title:        step.Step.Title,
		Description:  step.Step.Description,
		VideoURL:     step.Step.VideoURL,
		MarkComplete: step.MarkComplete,
	}
}

func (lr ListEnrolledStepResponse) FromListEntity(steps []entities.EnrolledStep) ListEnrolledStepResponse {
	data := ListEnrolledStepResponse{}

	for _, v := range steps {
		data = append(data, EnrolledStepResponse{}.FromEntity(v))
	}

	return data
}
