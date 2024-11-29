package entities

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
