package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type StepResponse struct {
	base.Base
	Title       string `json:"name"`
	Description string `json:"description"`
	VideoURL    string `json:"video_url"`
}

type ListStepResponse []StepResponse

func (r StepResponse) FromEntity(step entities.Step) StepResponse {
	return StepResponse{
		Base:        r.Base.FromEntity(step.Base),
		Title:       step.Title,
		Description: step.Description,
		VideoURL:    step.VideoURL,
	}
}

func (lr ListStepResponse) FromListEntity(steps []entities.Step) ListStepResponse {
	data := ListStepResponse{}

	for _, v := range steps {
		data = append(data, StepResponse{}.FromEntity(v))
	}

	return data
}
