package plant

import "github.com/labstack/echo/v4"

type PlantHandler interface {
	GetPlants(c echo.Context) error
	GetPlant(c echo.Context) error
	CreatePlant(c echo.Context) error
	UpdatePlant(c echo.Context) error
	DeletePlant(c echo.Context) error
}
