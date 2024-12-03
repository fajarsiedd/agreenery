package region

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/response"
	"go-agreenery/services/region"

	"github.com/labstack/echo/v4"
)

type regionHandler struct {
	service region.RegionService
}

func NewRegionHandler(service region.RegionService) *regionHandler {
	return &regionHandler{
		service: service,
	}
}

func (handler regionHandler) GetProvinces(c echo.Context) error {
	provinces, err := handler.service.GetProvinces()
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetProvincesSuccess, response.ListRegionResponse{}.FromListEntity(provinces))
}

func (handler regionHandler) GetRegencies(c echo.Context) error {
	provinceCode := c.Param("code")
	regencies, err := handler.service.GetRegencies(provinceCode)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetRegenciesSuccess, response.ListRegionResponse{}.FromListEntity(regencies))
}

func (handler regionHandler) GetDistricts(c echo.Context) error {
	regencyCode := c.Param("code")

	districts, err := handler.service.GetDistricts(regencyCode)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetDistrictsSuccess, response.ListRegionResponse{}.FromListEntity(districts))
}

func (handler regionHandler) GetVillages(c echo.Context) error {
	districtCode := c.Param("code")

	villages, err := handler.service.GetVillages(districtCode)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetVillagesSuccess, response.ListRegionResponse{}.FromListEntity(villages))
}
