package request

import (
	"go-agreenery/entities"
	"mime/multipart"
)

type WateringScheduleRequest struct {
	ID          string
	UserID      string
	PlantName   string         `form:"plant_name" validate:"required"`
	Image       multipart.File `form:"image"`
	RepeatEvery string         `form:"repeat_every" validate:"required"`
	StartDate   string         `form:"start_date" validate:"required"`
	EndDate     string         `form:"end_date" validate:"required"`
	TurnOnNotif bool           `form:"turn_on_notif"`
}

func (r WateringScheduleRequest) ToEntity() entities.WateringSchedule {
	return entities.WateringSchedule{
		Base:        entities.Base{ID: r.ID},
		UserID:      r.UserID,
		PlantName:   r.PlantName,
		ImageFile:   r.Image,
		RepeatEvery: r.RepeatEvery,
		TurnOnNotif: r.TurnOnNotif,
	}
}
