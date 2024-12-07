package models

import (
	"go-agreenery/constants/enums"
	"go-agreenery/entities"
)

type Category struct {
	Base
	Name   string
	Type   enums.CategoryType `gorm:"type:enum('article','plant','post')"`
	Plants []Plant
}

type ListCategory []Category

func (c Category) FromEntity(category entities.Category) Category {
	return Category{
		Base: c.Base.FromEntity(category.Base),
		Name: category.Name,
		Type: category.Type,
	}
}

func (c Category) ToEntity() entities.Category {
	return entities.Category{
		Base: c.Base.ToEntity(),
		Name: c.Name,
		Type: c.Type,
	}
}

func (lc ListCategory) FromListEntity(categories []entities.Category) ListCategory {
	listCategory := ListCategory{}

	for _, v := range categories {
		listCategory = append(listCategory, Category{}.FromEntity(v))
	}

	return listCategory
}

func (lc ListCategory) ToListEntity() []entities.Category {
	categories := []entities.Category{}

	for _, v := range lc {
		categories = append(categories, v.ToEntity())
	}

	return categories
}
