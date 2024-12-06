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

func (r CategoryRequest) ToEntity() entities.Category {
	return entities.Category{
		Base: entities.Base{ID: r.ID},
		Name: r.Name,
		Type: r.Type,
	}
}
