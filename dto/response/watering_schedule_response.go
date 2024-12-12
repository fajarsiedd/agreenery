package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
	"time"
)

type WateringScheduleResponse struct {
	base.Base
	PlantName   string          `json:"plant_name"`
	User        ProfileResponse `json:"user"`
	Image       string          `json:"image"`
	RepeatEvery string          `json:"repeat_every"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     time.Time       `json:"end_date"`
	TurnOnNotif bool            `json:"turn_on_notif"`
}

type ListWateringScheduleRsponse []WateringScheduleResponse

func (r WateringScheduleResponse) FromEntity(schedule entities.WateringSchedule) WateringScheduleResponse {
	return WateringScheduleResponse{
		Base:        r.Base.FromEntity(schedule.Base),
		PlantName:   schedule.PlantName,
		User:        r.User.FromEntity(schedule.User),
		Image:       schedule.Image,
		RepeatEvery: schedule.RepeatEvery,
		StartDate:   schedule.StartDate,
		EndDate:     schedule.EndDate,
		TurnOnNotif: schedule.TurnOnNotif,
	}
}

func (r ListWateringScheduleRsponse) FromListEntity(schedules []entities.WateringSchedule) ListWateringScheduleRsponse {
	data := ListWateringScheduleRsponse{}

	for _, v := range schedules {
		data = append(data, WateringScheduleResponse{}.FromEntity(v))
	}

	return data
}
