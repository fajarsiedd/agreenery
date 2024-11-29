package base

import (
	"go-agreenery/entities"
	"go-agreenery/helpers"

	"github.com/labstack/echo/v4"
)

type Meta struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
	TotalItems int `json:"total_items"`
}

type ResponseWrapper struct {
	Meta Meta `json:"meta"`
	Data any  `json:"data,omitempty"`
}

type ResponsePaginationWrapper struct {
	Meta       Meta       `json:"meta"`
	Data       any        `json:"data,omitempty"`
	Pagination Pagination `json:"pagination"`
}

func SuccessResponse(c echo.Context, message string, data any) error {
	statusCode := helpers.GetStatusCodeBySuccessMessage(message)

	return c.JSON(statusCode, ResponseWrapper{
		Meta: Meta{
			Status:  true,
			Code:    statusCode,
			Message: message,
		},
		Data: data,
	})
}

func SuccessResponsePagination(c echo.Context, message string, pagination entities.Pagination, data any) error {
	statusCode := helpers.GetStatusCodeBySuccessMessage(message)

	return c.JSON(statusCode, ResponsePaginationWrapper{
		Meta: Meta{
			Status:  true,
			Code:    statusCode,
			Message: message,
		},
		Data: data,
		Pagination: Pagination{
			Page:       pagination.Page,
			Limit:      pagination.Limit,
			TotalPages: pagination.TotalPages,
			TotalItems: pagination.TotalItems,
		},
	})
}

func ErrorResponse(c echo.Context, err error) error {
	statusCode := helpers.GetStatusCodeByErr(err)

	return c.JSON(statusCode, ResponseWrapper{
		Meta: Meta{
			Status: false,
			Code:   statusCode,
			Error:  err.Error(),
		},
	})
}
