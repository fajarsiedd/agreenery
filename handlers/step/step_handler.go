package step

import (
	"github.com/labstack/echo/v4"
)

type StepHandler interface {
	CreateStep(c echo.Context) error
	UpdateStep(c echo.Context) error
	DeleteStep(c echo.Context) error
}
