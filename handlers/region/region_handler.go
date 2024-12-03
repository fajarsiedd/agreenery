package region

import "github.com/labstack/echo/v4"

type RegionHandler interface {
	GetProvinces(c echo.Context) error
	GetRegencies(c echo.Context) error
	GetDistricts(c echo.Context) error
	GetVillages(c echo.Context) error
}
