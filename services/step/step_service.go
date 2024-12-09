package step

import "go-agreenery/entities"

type StepService interface {
	CreateStep(step entities.Step) (entities.Plant, error)
	UpdateStep(step entities.Step) (entities.Plant, error)
	DeleteStep(id string) (entities.Plant, error)
}
