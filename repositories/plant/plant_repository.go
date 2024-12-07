package plant

import "go-agreenery/entities"

type PlantRepository interface {
	GetPlants(filter entities.Filter) ([]entities.Plant, entities.Pagination, error)
	GetPlant(id string) (entities.Plant, error)
	CreatePlant(category entities.Plant) (entities.Plant, error)
	UpdatePlant(category entities.Plant) (entities.Plant, error)
	DeletePlant(id string) error
}
