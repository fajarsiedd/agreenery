package step

import (
	"go-agreenery/entities"
	"go-agreenery/repositories/step"
)

type stepService struct {
	repository step.StepRepository
}

func NewStepService(r step.StepRepository) *stepService {
	return &stepService{
		repository: r,
	}
}

func (s stepService) CreateStep(step entities.Step) (entities.Plant, error) {
	return s.repository.CreateStep(step)
}

func (s stepService) UpdateStep(step entities.Step) (entities.Plant, error) {
	return s.repository.UpdateStep(step)
}

func (s stepService) DeleteStep(id string) (entities.Plant, error) {
	return s.repository.DeleteStep(id)
}
