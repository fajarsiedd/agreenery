package models

import "go-agreenery/entities"

type Plant struct {
	Base
	Name        string
	Description string
	Image       string
	CategoryID  string   `gorm:"size:191"`
	Category    Category `gorm:"foreignKey:CategoryID;references:ID"`
}

type ListPlant []Plant

func (p Plant) FromEntity(plant entities.Plant) Plant {
	return Plant{
		Base:        p.Base.FromEntity(plant.Base),
		Name:        plant.Name,
		Description: plant.Description,
		Image:       plant.Image,
		CategoryID:  plant.CategoryID,
		Category:    p.Category.FromEntity(plant.Category),
	}
}

func (p Plant) ToEntity() entities.Plant {
	return entities.Plant{
		Base:        p.Base.ToEntity(),
		Name:        p.Name,
		Description: p.Description,
		Image:       p.Image,
		CategoryID:  p.CategoryID,
		Category:    p.Category.ToEntity(),
	}
}

func (lp ListPlant) FromListEntity(categories []entities.Plant) ListPlant {
	data := ListPlant{}

	for _, v := range categories {
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
