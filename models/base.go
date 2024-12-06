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

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New().String()
	return
}

func (b Base) FromEntity(base entities.Base) Base {
	return Base{
		ID:        base.ID,
		CreatedAt: base.CreatedAt,
		UpdatedAt: base.UpdatedAt,
		DeletedAt: base.DeletedAt,
	}
}

func (b Base) ToEntity() entities.Base {
	return entities.Base{
		ID:        b.ID,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
		DeletedAt: b.DeletedAt,
	}
}
