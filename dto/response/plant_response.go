package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type PlantResponse struct {
	base.Base
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Image        string           `json:"image"`
	Fertilizer   string           `json:"fertilizer"`
	PlantingTips string           `json:"planting_tips"`
	Category     string           `json:"category"`
	Steps        ListStepResponse `json:"steps,omitempty"`
}

type ListPlantResponse []PlantResponse

func (r PlantResponse) FromEntity(plant entities.Plant) PlantResponse {
	return PlantResponse{
		Base:         r.Base.FromEntity(plant.Base),
		Name:         plant.Name,
		Description:  plant.Description,
		Image:        plant.Image,
		Fertilizer:   plant.Fertilizer,
		PlantingTips: plant.PlantingTips,
		Category:     plant.Category.Name,
		Steps:        r.Steps.FromListEntity(plant.Steps),
	}
}

func (lr ListPlantResponse) FromListEntity(plants []entities.Plant) ListPlantResponse {
	data := ListPlantResponse{}

	for _, v := range plants {
		data = append(data, PlantResponse{}.FromEntity(v))
	}

	return data
}
