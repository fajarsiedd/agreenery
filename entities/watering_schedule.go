package entities

import (
	"mime/multipart"
	"time"
)

type WateringSchedule struct {
	Base
	UserID      string
	User        User
	PlantName   string
	Image       string
	ImageFile   multipart.File
	RepeatEvery string
	StartDate   time.Time
	EndDate     time.Time
	TurnOnNotif bool
}
