package notification

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/request"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"go-agreenery/middlewares"
	"go-agreenery/services/notification"
	"time"

	"github.com/labstack/echo/v4"
)

type notificationHandler struct {
	service notification.NotificationService
}

func NewNotificationHandler(s notification.NotificationService) *notificationHandler {
	return &notificationHandler{
		service: s,
	}
}

func (h notificationHandler) GetNotifications(c echo.Context) error {
	filter, err := helpers.GetFilter(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	result, pagination, err := h.service.GetNotifications(filter)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponsePagination(c, constants.GetNotificationsSuccess, pagination, response.ListNotificationResponse{}.FromListEntity(result))
}

func (h notificationHandler) CreateNotification(c echo.Context) error {
	req := request.NotificationRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	notification := req.ToEntity()

	location, _ := time.LoadLocation("Asia/Bangkok")
	parsedDate, err := time.Parse("2006-01-02", req.SendAt)
	if err != nil {
		return base.ErrorResponse(c, constants.ErrInvalidSendAt)
	}
	notification.SendAt = time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, location)

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	notification.UserID = claims.UserID

	result, err := h.service.CreateNotification(notification)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.CreateNotificationSuccess, response.NotificationResponse{}.FromEntity(result))
}

func (h notificationHandler) UpdateNotification(c echo.Context) error {
	id := c.Param("id")

	req := request.NotificationRequest{ID: id}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	notification := req.ToEntity()

	location, _ := time.LoadLocation("Asia/Bangkok")
	parsedDate, err := time.Parse("2006-01-02", req.SendAt)
	if err != nil {
		return base.ErrorResponse(c, constants.ErrInvalidSendAt)
	}
	notification.SendAt = time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, location)

	result, err := h.service.UpdateNotification(notification)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.UpdateNotificationSuccess, response.NotificationResponse{}.FromEntity(result))
}

func (h notificationHandler) DeleteNotification(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteNotification(id); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.DeleteNotificationSuccess, nil)
}

func (h notificationHandler) SendNotification(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.SendNotification(id); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.SendNotificationSuccess, nil)
}
