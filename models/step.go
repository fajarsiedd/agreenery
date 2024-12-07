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

func (p Step) FromEntity(step entities.Step) Step {
	return Step{
		Base:        p.Base.FromEntity(step.Base),
		Title:       step.Title,
		Description: step.Description,
		VideoURL:    step.VideoURL,
		PlantID:     step.PlantID,
	}
}

func (p Step) ToEntity() entities.Step {
	return entities.Step{
		Base:        p.Base.ToEntity(),
		Title:       p.Title,
		Description: p.Description,
		VideoURL:    p.VideoURL,
		PlantID:     p.PlantID,
	}
}

func (lp ListStep) FromListEntity(categories []entities.Step) ListStep {
	data := ListStep{}

	for _, v := range categories {
		data = append(data, Step{}.FromEntity(v))
	}

	return data
}

func (lp ListStep) ToListEntity() []entities.Step {
	data := []entities.Step{}

	for _, v := range lp {
		data = append(data, v.ToEntity())
	}

	return data
}
