package models

import "go-agreenery/entities"

type Step struct {
	Base
	Title       string
	Description string
	VideoURL    string
	PlantID     string `gorm:"size:191"`
}

type ListStep []Step

func (s Step) FromEntity(step entities.Step) Step {
	return Step{
		Base:        s.Base.FromEntity(step.Base),
		Title:       step.Title,
		Description: step.Description,
		VideoURL:    step.VideoURL,
		PlantID:     step.PlantID,
	}
}

func (s Step) ToEntity() entities.Step {
	return entities.Step{
		Base:        s.Base.ToEntity(),
		Title:       s.Title,
		Description: s.Description,
		VideoURL:    s.VideoURL,
		PlantID:     s.PlantID,
	}
}

func (ls ListStep) FromListEntity(steps []entities.Step) ListStep {
	data := ListStep{}

	for _, v := range steps {
		data = append(data, Step{}.FromEntity(v))
	}

	return data
}

func (ls ListStep) ToListEntity() []entities.Step {
	data := []entities.Step{}

	for _, v := range ls {
		data = append(data, v.ToEntity())
	}

	return data
}
