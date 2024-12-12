package models

import (
	"go-agreenery/entities"
	"time"
)

type WateringSchedule struct {
	Base
	UserID      string `gorm:"size:191"`
	User        User   `gorm:"foreignKey:UserID;references:ID"`
	PlantName   string
	Image       string
	RepeatEvery string
	StartDate   time.Time
	EndDate     time.Time
	TurnOnNotif bool `gorm:"default:false"`
}

type ListWateringSchedule []WateringSchedule

func (s WateringSchedule) FromEntity(schedule entities.WateringSchedule) WateringSchedule {
	return WateringSchedule{
		Base:        s.Base.FromEntity(schedule.Base),
		UserID:      schedule.UserID,
		User:        s.User.FromEntity(schedule.User),
		PlantName:   schedule.PlantName,
		Image:       schedule.Image,
		RepeatEvery: schedule.RepeatEvery,
		StartDate:   schedule.StartDate,
		EndDate:     schedule.EndDate,
		TurnOnNotif: schedule.TurnOnNotif,
	}
}

func (s WateringSchedule) ToEntity() entities.WateringSchedule {
	return entities.WateringSchedule{
		Base:        s.Base.ToEntity(),
		UserID:      s.UserID,
		User:        s.User.ToEntity(),
		PlantName:   s.PlantName,
		Image:       s.Image,
		RepeatEvery: s.RepeatEvery,
		StartDate:   s.StartDate,
		EndDate:     s.EndDate,
		TurnOnNotif: s.TurnOnNotif,
	}
}

func (ls ListWateringSchedule) FromListEntity(schedules []entities.WateringSchedule) ListWateringSchedule {
	data := ListWateringSchedule{}

	for _, v := range schedules {
		data = append(data, WateringSchedule{}.FromEntity(v))
	}

	return data
}

func (ls ListWateringSchedule) ToListEntity() []entities.WateringSchedule {
	data := []entities.WateringSchedule{}

	for _, v := range ls {
		data = append(data, v.ToEntity())
	}

	return data
}
