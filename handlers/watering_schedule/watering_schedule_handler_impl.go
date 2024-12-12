package watering_schedule

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/request"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"go-agreenery/middlewares"
	"go-agreenery/services/watering_schedule"
	"io"
	"time"

	"github.com/h2non/filetype"
	"github.com/labstack/echo/v4"
)

type wateringScheduleHandler struct {
	service watering_schedule.WateringScheduleService
}

func NewWateringScheduleHandler(s watering_schedule.WateringScheduleService) *wateringScheduleHandler {
	return &wateringScheduleHandler{
		service: s,
	}
}

func (h wateringScheduleHandler) GetWateringSchedules(c echo.Context) error {
	filter, err := helpers.GetFilter(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	result, pagination, err := h.service.GetWateringSchedules(filter)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponsePagination(c, constants.GetWateringSchedulesSuccess, pagination, response.ListWateringScheduleRsponse{}.FromListEntity(result))
}

func (h wateringScheduleHandler) GetWateringSchedule(c echo.Context) error {
	id := c.Param("id")

	result, err := h.service.GetWateringSchedule(id)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetWateringScheduleSuccess, response.WateringScheduleResponse{}.FromEntity(result))
}

func (h wateringScheduleHandler) CreateWateringSchedule(c echo.Context) error {
	req := request.WateringScheduleRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	file, err := c.FormFile("image")
	if err == nil {
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
	}

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	schedule := req.ToEntity()
	schedule.UserID = claims.UserID

	location, _ := time.LoadLocation("Asia/Jakarta")
	parsedStartDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return base.ErrorResponse(c, constants.ErrInvalidStartDateParam)
	}
	schedule.StartDate = time.Date(parsedStartDate.Year(), parsedStartDate.Month(), parsedStartDate.Day(), 0, 0, 0, 0, location)

	parsedEndDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return base.ErrorResponse(c, constants.ErrInvalidEndDateParam)
	}
	schedule.EndDate = time.Date(parsedEndDate.Year(), parsedEndDate.Month(), parsedEndDate.Day(), 0, 0, 0, 0, location)

	result, err := h.service.CreateWateringSchedule(schedule)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.CreateWateringScheduleSuccess, response.WateringScheduleResponse{}.FromEntity(result))
}

func (h wateringScheduleHandler) UpdateWateringSchedule(c echo.Context) error {
	id := c.Param("id")

	req := request.WateringScheduleRequest{ID: id}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	file, err := c.FormFile("image")
	if err == nil {
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
		if !filetype.IsImage(buf) && !filetype.IsVideo(buf) {
			return base.ErrorResponse(c, constants.ErrOnlyImageAndVideoAllowed)
		}

		req.Image = blobFile
	}

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	schedule := req.ToEntity()

	location, _ := time.LoadLocation("Asia/Jakarta")
	parsedStartDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return base.ErrorResponse(c, constants.ErrInvalidStartDateParam)
	}
	schedule.StartDate = time.Date(parsedStartDate.Year(), parsedStartDate.Month(), parsedStartDate.Day(), 0, 0, 0, 0, location)

	parsedEndDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		return base.ErrorResponse(c, constants.ErrInvalidEndDateParam)
	}
	schedule.EndDate = time.Date(parsedEndDate.Year(), parsedEndDate.Month(), parsedEndDate.Day(), 0, 0, 0, 0, location)

	result, err := h.service.UpdateWateringSchedule(schedule, claims.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.UpdateWateringScheduleSuccess, response.WateringScheduleResponse{}.FromEntity(result))
}

func (h wateringScheduleHandler) DeleteWateringSchedule(c echo.Context) error {
	id := c.Param("id")

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := h.service.DeleteWateringSchedule(id, claims.UserID); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.DeleteWateringScheduleSuccess, nil)
}
