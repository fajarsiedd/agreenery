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

func NewRegionHandler(s region.RegionService) *regionHandler {
	return &regionHandler{
		service: s,
	}
}

func (h regionHandler) GetProvinces(c echo.Context) error {
	provinces, err := h.service.GetProvinces()
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetProvincesSuccess, response.ListRegionResponse{}.FromListEntity(provinces))
}

func (h regionHandler) GetRegencies(c echo.Context) error {
	provinceCode := c.Param("code")
	regencies, err := h.service.GetRegencies(provinceCode)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetRegenciesSuccess, response.ListRegionResponse{}.FromListEntity(regencies))
}

func (h regionHandler) GetDistricts(c echo.Context) error {
	regencyCode := c.Param("code")

	districts, err := h.service.GetDistricts(regencyCode)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetDistrictsSuccess, response.ListRegionResponse{}.FromListEntity(districts))
}

func (h regionHandler) GetVillages(c echo.Context) error {
	districtCode := c.Param("code")

	villages, err := h.service.GetVillages(districtCode)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetVillagesSuccess, response.ListRegionResponse{}.FromListEntity(villages))
}
