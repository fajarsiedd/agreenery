package models

import "go-agreenery/entities"

type EnrolledStep struct {
	Base
	UserID          string `gorm:"size:191"`
	EnrolledPlantID string `gorm:"size:191"`
	StepID          string `gorm:"size:191"`
	Step            Step   `gorm:"foreignKey:StepID;references:ID"`
	MarkComplete    bool   `gorm:"default:false"`
}

type ListEnrolledStep []EnrolledStep

func (s EnrolledStep) FromEntity(step entities.EnrolledStep) EnrolledStep {
	return EnrolledStep{
		Base:            s.Base.FromEntity(step.Base),
		UserID:          step.UserID,
		EnrolledPlantID: step.EnrolledPlantID,
		StepID:          step.StepID,
		Step:            s.Step.FromEntity(step.Step),
		MarkComplete:    step.MarkComplete,
	}
}

func (p EnrolledStep) ToEntity() entities.EnrolledStep {
	return entities.EnrolledStep{
		Base:            p.Base.ToEntity(),
		UserID:          p.UserID,
		EnrolledPlantID: p.EnrolledPlantID,
		StepID:          p.StepID,
		Step:            p.Step.ToEntity(),
		MarkComplete:    p.MarkComplete,
	}
}

func (lp ListEnrolledStep) FromListEntity(steps []entities.EnrolledStep) ListEnrolledStep {
	data := ListEnrolledStep{}

	for _, v := range steps {
		data = append(data, EnrolledStep{}.FromEntity(v))
	}

	return data
}

func (lp ListEnrolledStep) ToListEntity() []entities.EnrolledStep {
	data := []entities.EnrolledStep{}

	for _, v := range lp {
		data = append(data, v.ToEntity())
	}

	return data
}
