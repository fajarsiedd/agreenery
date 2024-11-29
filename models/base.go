package models

import (
	"go-agreenery/entities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.ID = uuid.New().String()
	return
}

func (Base) FromEntity(baseEntity entities.Base) Base {
	return Base{
		ID:        baseEntity.ID,
		CreatedAt: baseEntity.CreatedAt,
		UpdatedAt: baseEntity.UpdatedAt,
		DeletedAt: baseEntity.DeletedAt,
	}
}

func (base Base) ToEntity() entities.Base {
	return entities.Base{
		ID:        base.ID,
		CreatedAt: base.CreatedAt,
		UpdatedAt: base.UpdatedAt,
		DeletedAt: base.DeletedAt,
	}
}
