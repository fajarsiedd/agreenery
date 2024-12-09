package plant

import "go-agreenery/entities"

type PlantRepository interface {
	GetPlants(filter entities.Filter) ([]entities.Plant, entities.Pagination, error)
	GetPlant(id string) (entities.Plant, error)
	CreatePlant(plant entities.Plant) (entities.Plant, error)
	UpdatePlant(plant entities.Plant) (entities.Plant, error)
	DeletePlant(id string) error
}
