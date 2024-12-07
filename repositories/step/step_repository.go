package step

import "go-agreenery/entities"

type StepRepository interface {
	CreateStep(category entities.Step) (entities.Plant, error)
	UpdateStep(category entities.Step) (entities.Plant, error)
	DeleteStep(id string) (entities.Plant, error)
}
