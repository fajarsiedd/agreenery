package watering_schedule

import "github.com/labstack/echo/v4"

type WateringScheduleHandler interface {
	GetWateringSchedules(c echo.Context) error
	GetWateringSchedule(c echo.Context) error
	CreateWateringSchedule(c echo.Context) error
	UpdateWateringSchedule(c echo.Context) error
	DeleteWateringSchedule(c echo.Context) error
}
