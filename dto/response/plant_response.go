package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type PlantResponse struct {
	base.Base
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Category    string `json:"category"`
}

type ListPlantResponse []PlantResponse

func (p PlantResponse) FromEntity(plant entities.Plant) PlantResponse {
	return PlantResponse{
		Base:        p.Base.FromEntity(plant.Base),
		Name:        plant.Name,
		Description: plant.Description,
		Image:       plant.Image,
		Category:    plant.Category.Name,
	}
}

func (lp ListPlantResponse) FromListEntity(plants []entities.Plant) ListPlantResponse {
	data := ListPlantResponse{}

	for _, v := range plants {
		data = append(data, PlantResponse{}.FromEntity(v))
	}

	return data
}
