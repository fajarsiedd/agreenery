package enrollment

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/request"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"go-agreenery/middlewares"
	"go-agreenery/services/enrollment"

	"github.com/labstack/echo/v4"
)

type enrollmentHandler struct {
	service enrollment.EnrollmentService
}

func NewEnrollmentHandler(s enrollment.EnrollmentService) *enrollmentHandler {
	return &enrollmentHandler{
		service: s,
	}
}

func (h enrollmentHandler) CreateEnrollment(c echo.Context) error {
	req := request.EnrollmentRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	req.UserID = claims.UserID

	plant := req.ToEntity()

	result, err := h.service.CreateEnrollment(plant)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.CreateEnrollmentSuccess, response.EnrolledPlantResponse{}.FromEntity(result))
}

func (h enrollmentHandler) GetEnrollments(c echo.Context) error {
	filter, err := helpers.GetFilter(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	filter.UserID = claims.UserID

	result, pagination, err := h.service.GetEnrollments(filter)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponsePagination(c, constants.GetEnrollmentsSuccess, pagination, response.ListEnrolledPlantResponse{}.FromListEntity(result))
}

func (h enrollmentHandler) GetEnrollment(c echo.Context) error {
	id := c.Param("enrollmentID")

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	result, err := h.service.GetEnrollment(id, claims.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetEnrollmentSuccess, response.EnrolledPlantResponse{}.FromEntity(result))
}

func (h enrollmentHandler) MarkStepAsComplete(c echo.Context) error {
	id := c.Param("stepID")

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	result, err := h.service.MarkStepAsComplete(id, claims.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.MarkStepAsCompleteSuccess, response.EnrolledPlantResponse{}.FromEntity(result))
}

func (h enrollmentHandler) SetEnrollmentStatusAsDone(c echo.Context) error {
	id := c.Param("enrollmentID")

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	result, err := h.service.SetEnrollmentStatusAsDone(id, claims.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.SetEnrollmentStatusAsDoneSuccess, response.EnrolledPlantResponse{}.FromEntity(result))
}

func (h enrollmentHandler) DeleteEnrollment(c echo.Context) error {
	id := c.Param("enrollmentID")

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	err = h.service.DeleteEnrollment(id, claims.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.DeleteEnrollmentSuccess, nil)
}
