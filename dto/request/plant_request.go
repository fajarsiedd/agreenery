package request

import (
	"go-agreenery/entities"
	"mime/multipart"
)

type PlantRequest struct {
	ID          string
	Name        string         `form:"name" validate:"required"`
	Description string         `form:"description" validate:"required"`
	Image       multipart.File `form:"image"`
	CategoryID  string         `form:"category_id" validate:"required"`
}

func (r PlantRequest) ToEntity() entities.Plant {
	return entities.Plant{
		Base:        entities.Base{ID: r.ID},
		Name:        r.Name,
		Description: r.Description,
		ImageFile:   r.Image,
		CategoryID:  r.CategoryID,
	}
}
