package request

import (
	"go-agreenery/constants/enums"
	"go-agreenery/entities"
)

type CategoryRequest struct {
	ID   string
	Name string             `json:"name" validate:"required"`
	Type enums.CategoryType `json:"type" validate:"required,oneof=article post plant"`
}

func (c CategoryRequest) ToEntity() entities.Category {
	return entities.Category{
		Base: entities.Base{ID: c.ID},
		Name: c.Name,
		Type: c.Type,
	}
}
