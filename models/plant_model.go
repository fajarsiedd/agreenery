package models

import "go-agreenery/entities"

type Plant struct {
	Base
	Name         string
	Description  string
	Image        string
	Fertilizer   string
	PlantingTips string
	CategoryID   string   `gorm:"size:191"`
	Category     Category `gorm:"foreignKey:CategoryID;references:ID"`
	Steps        ListStep
}

type ListPlant []Plant

func (p Plant) FromEntity(plant entities.Plant) Plant {
	return Plant{
		Base:         p.Base.FromEntity(plant.Base),
		Name:         plant.Name,
		Description:  plant.Description,
		Fertilizer:   plant.Fertilizer,
		PlantingTips: plant.PlantingTips,
		Image:        plant.Image,
		CategoryID:   plant.CategoryID,
		Category:     p.Category.FromEntity(plant.Category),
		Steps:        p.Steps.FromListEntity(plant.Steps),
	}
}

func (p Plant) ToEntity() entities.Plant {
	return entities.Plant{
		Base:         p.Base.ToEntity(),
		Name:         p.Name,
		Description:  p.Description,
		Fertilizer:   p.Fertilizer,
		PlantingTips: p.PlantingTips,
		Image:        p.Image,
		CategoryID:   p.CategoryID,
		Category:     p.Category.ToEntity(),
		Steps:        p.Steps.ToListEntity(),
	}
}

func (lp ListPlant) FromListEntity(plants []entities.Plant) ListPlant {
	data := ListPlant{}

	for _, v := range plants {
		data = append(data, Plant{}.FromEntity(v))
	}

	return data
}

func (lp ListPlant) ToListEntity() []entities.Plant {
	data := []entities.Plant{}

	for _, v := range lp {
		data = append(data, v.ToEntity())
	}

	return data
}
