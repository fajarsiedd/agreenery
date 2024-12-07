package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type EnrolledPlantResponse struct {
	base.Base
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Image       string                   `json:"image"`
	Category    string                   `json:"category"`
	Steps       ListEnrolledStepResponse `json:"steps,omitempty"`
	IsDone      bool                     `json:"is_done"`
	UserID      string                   `json:"user_id"`
}

type ListEnrolledPlantResponse []EnrolledPlantResponse

func (r EnrolledPlantResponse) FromEntity(plant entities.EnrolledPlant) EnrolledPlantResponse {
	return EnrolledPlantResponse{
		Base:        r.Base.FromEntity(plant.Base),
		Name:        plant.Plant.Name,
		Description: plant.Plant.Description,
		Image:       plant.Plant.Image,
		Category:    plant.Plant.Category.Name,
		Steps:       r.Steps.FromListEntity(plant.EnrolledSteps),
		IsDone:      plant.IsDone,
		UserID:      plant.UserID,
	}
}

func (lr ListEnrolledPlantResponse) FromListEntity(plants []entities.EnrolledPlant) ListEnrolledPlantResponse {
	data := ListEnrolledPlantResponse{}

	for _, v := range plants {
		data = append(data, EnrolledPlantResponse{}.FromEntity(v))
	}

	return data
}
