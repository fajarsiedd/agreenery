package watering_schedule

import "go-agreenery/entities"

type WateringScheduleRepository interface {
	GetWateringSchedules(filter entities.Filter) ([]entities.WateringSchedule, entities.Pagination, error)
	GetWateringSchedule(id string) (entities.WateringSchedule, error)
	CreateWateringSchedule(schedule entities.WateringSchedule) (entities.WateringSchedule, error)
	UpdateWateringSchedule(schedule entities.WateringSchedule, currUserID string) (entities.WateringSchedule, error)
	DeleteWateringSchedule(id, currUserID string) (string, error)
}
