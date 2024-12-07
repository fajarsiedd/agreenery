package step

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/request"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"go-agreenery/services/step"

	"github.com/labstack/echo/v4"
)

type stepHandler struct {
	service step.StepService
}

func NewStepHandler(s step.StepService) *stepHandler {
	return &stepHandler{
		service: s,
	}
}

func (h stepHandler) CreateStep(c echo.Context) error {
	req := request.StepRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	step := req.ToEntity()

	result, err := h.service.CreateStep(step)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.CreateStepSuccess, response.PlantResponse{}.FromEntity(result))
}

func (h stepHandler) UpdateStep(c echo.Context) error {
	id := c.Param("id")

	req := request.StepRequest{ID: id}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	step := req.ToEntity()

	result, err := h.service.UpdateStep(step)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.UpdateStepSuccess, response.PlantResponse{}.FromEntity(result))
}

func (h stepHandler) DeleteStep(c echo.Context) error {
	id := c.Param("id")

	result, err := h.service.DeleteStep(id)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.DeleteStepSuccess, response.PlantResponse{}.FromEntity(result))
}
