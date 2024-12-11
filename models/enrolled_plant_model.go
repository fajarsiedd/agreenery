package models

import "go-agreenery/entities"

type EnrolledPlant struct {
	Base
	UserID        string `gorm:"size:191"`
	PlantID       string `gorm:"size:191"`
	Plant         Plant  `gorm:"foreignKey:PlantID;references:ID"`
	IsDone        bool   `gorm:"default:false"`
	EnrolledSteps ListEnrolledStep
}

type ListEnrolledPlant []EnrolledPlant

func (p EnrolledPlant) FromEntity(plant entities.EnrolledPlant) EnrolledPlant {
	return EnrolledPlant{
		Base:          p.Base.FromEntity(plant.Base),
		UserID:        plant.UserID,
		PlantID:       plant.PlantID,
		Plant:         p.Plant.FromEntity(plant.Plant),
		EnrolledSteps: p.EnrolledSteps.FromListEntity(plant.EnrolledSteps),
		IsDone:        plant.IsDone,
	}
}

func (p EnrolledPlant) ToEntity() entities.EnrolledPlant {
	return entities.EnrolledPlant{
		Base:          p.Base.ToEntity(),
		UserID:        p.UserID,
		PlantID:       p.PlantID,
		Plant:         p.Plant.ToEntity(),
		EnrolledSteps: p.EnrolledSteps.ToListEntity(),
		IsDone:        p.IsDone,
	}
}

func (lp ListEnrolledPlant) FromListEntity(plants []entities.EnrolledPlant) ListEnrolledPlant {
	data := ListEnrolledPlant{}

	for _, v := range plants {
		data = append(data, EnrolledPlant{}.FromEntity(v))
	}

	return data
}

func (lp ListEnrolledPlant) ToListEntity() []entities.EnrolledPlant {
	data := []entities.EnrolledPlant{}

	for _, v := range lp {
		data = append(data, v.ToEntity())
	}

	return data
}
