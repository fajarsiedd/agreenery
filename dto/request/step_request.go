package request

import (
	"go-agreenery/entities"
)

type StepRequest struct {
	ID          string
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	VideoURL    string `json:"video_url" validate:"required"`
	PlantID     string `json:"plant_id" validate:"required"`
}

func (r StepRequest) ToEntity() entities.Step {
	return entities.Step{
		Base:        entities.Base{ID: r.ID},
		Title:       r.Title,
		Description: r.Description,
		VideoURL:    r.VideoURL,
		PlantID:     r.PlantID,
	}
}
