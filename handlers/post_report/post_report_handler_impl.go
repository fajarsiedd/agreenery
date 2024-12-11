package post_report

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/request"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"go-agreenery/middlewares"
	"go-agreenery/services/post_report"

	"github.com/labstack/echo/v4"
)

type postReportHandler struct {
	service post_report.PostReportService
}

func NewPostReportHandler(s post_report.PostReportService) *postReportHandler {
	return &postReportHandler{
		service: s,
	}
}

func (h postReportHandler) GetPostReports(c echo.Context) error {
	filter, err := helpers.GetFilter(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	filter.UserID = claims.UserID

	result, pagination, err := h.service.GetPostReports(filter)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponsePagination(c, constants.GetPostReportsSuccess, pagination, response.ListPostReportResponse{}.FromListEntity(result))
}

func (h postReportHandler) CreatePostReport(c echo.Context) error {
	req := request.PostReportRequest{}

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

	postReport := req.ToEntity()
	postReport.UserID = claims.UserID

	result, err := h.service.CreatePostReport(postReport)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.CreatePostReportSuccess, response.PostReportResponse{}.FromEntity(result))
}

func (h postReportHandler) DeletePostReport(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeletePostReport(id); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.DeletePostReportSuccess, nil)
}

func (h postReportHandler) SendWarning(c echo.Context) error {
	id := c.Param("id")
	req := request.PostReportActionRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	postReport := req.ToEntity()
	postReport.ID = id

	if err := h.service.SendWarning(postReport); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.SendWarningSuccess, nil)
}

func (h postReportHandler) DeletePostWithMessage(c echo.Context) error {
	id := c.Param("id")
	req := request.PostReportActionRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	postReport := req.ToEntity()
	postReport.ID = id

	if err := h.service.DeletePostWithMessage(postReport); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.DeletePostSuccess, nil)
}
