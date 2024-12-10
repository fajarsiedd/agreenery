package models

import (
	"go-agreenery/constants/enums"
	"go-agreenery/entities"
)

type Category struct {
	Base
	Name       string
	Type       enums.CategoryType `gorm:"type:enum('article','plant','post')"`
	Plants     ListPlant
	Article    ListArticle
	Posts      ListPost
	CountPosts int64 `gorm:"<-:false;-:migration"`
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
		Base:       c.Base.ToEntity(),
		Name:       c.Name,
		Type:       c.Type,
		CountPosts: c.CountPosts,
	}
}

func (lc ListCategory) FromListEntity(categories []entities.Category) ListCategory {
	data := ListCategory{}

	for _, v := range categories {
		data = append(data, Category{}.FromEntity(v))
	}

	return data
}

func (lc ListCategory) ToListEntity() []entities.Category {
	data := []entities.Category{}

	for _, v := range lc {
		data = append(data, v.ToEntity())
	}

	return data
}
