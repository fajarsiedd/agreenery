package base

import (
	"go-agreenery/entities"
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (Base) FromEntity(baseEntity entities.Base) Base {
	return Base{
		ID:        baseEntity.ID,
		CreatedAt: baseEntity.CreatedAt,
		UpdatedAt: baseEntity.UpdatedAt,
		DeletedAt: baseEntity.DeletedAt,
	}
}
