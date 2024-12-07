package plant

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/request"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"go-agreenery/services/plant"
	"io"

	"github.com/h2non/filetype"
	"github.com/labstack/echo/v4"
)

type plantHandler struct {
	service plant.PlantService
}

func NewPlantHandler(s plant.PlantService) *plantHandler {
	return &plantHandler{
		service: s,
	}
}

func (h plantHandler) GetPlants(c echo.Context) error {
	filter, err := helpers.GetFilter(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	result, pagination, err := h.service.GetPlants(filter)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponsePagination(c, constants.GetPlantsSuccess, pagination, response.ListPlantResponse{}.FromListEntity(result))
}

func (h plantHandler) GetPlant(c echo.Context) error {
	id := c.Param("id")

	result, err := h.service.GetPlant(id)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetPlantSuccess, response.PlantResponse{}.FromEntity(result))
}

func (h plantHandler) CreatePlant(c echo.Context) error {
	req := request.PlantRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	file, err := c.FormFile("image")
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	var maxFileSize int64 = 1048576 * 2
	if file.Size > maxFileSize {
		return base.ErrorResponse(c, constants.ErrFileSizeExceedsLimit)
	}

	blobFile, err := file.Open()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	defer blobFile.Close()

	temp, _ := file.Open()
	buf, _ := io.ReadAll(temp)
	if !filetype.IsImage(buf) {
		return base.ErrorResponse(c, constants.ErrOnlyImageAllowed)
	}

	req.Image = blobFile
	plant := req.ToEntity()

	result, err := h.service.CreatePlant(plant)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.CreatePlantSuccess, response.PlantResponse{}.FromEntity(result))
}

func (h plantHandler) UpdatePlant(c echo.Context) error {
	id := c.Param("id")

	req := request.PlantRequest{ID: id}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	file, err := c.FormFile("image")
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	var maxFileSize int64 = 1048576 * 2
	if file.Size > maxFileSize {
		return base.ErrorResponse(c, constants.ErrFileSizeExceedsLimit)
	}

	blobFile, err := file.Open()
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	defer blobFile.Close()

	temp, _ := file.Open()
	buf, _ := io.ReadAll(temp)
	if !filetype.IsImage(buf) {
		return base.ErrorResponse(c, constants.ErrOnlyImageAllowed)
	}

	req.Image = blobFile
	plant := req.ToEntity()

	result, err := h.service.UpdatePlant(plant)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.UpdatePlantSuccess, response.PlantResponse{}.FromEntity(result))
}

func (h plantHandler) DeletePlant(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeletePlant(id); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.DeletePlantSuccess, nil)
}
