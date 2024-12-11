package user_notification

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"go-agreenery/middlewares"
	"go-agreenery/services/user_notification"

	"github.com/labstack/echo/v4"
)

type userNotificationHandler struct {
	service user_notification.UserNotificationService
}

func NewUserNotificationHandler(s user_notification.UserNotificationService) *userNotificationHandler {
	return &userNotificationHandler{
		service: s,
	}
}

func (h userNotificationHandler) GetUserNotifications(c echo.Context) error {
	filter, err := helpers.GetFilter(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	filter.UserID = claims.UserID

	result, pagination, err := h.service.GetUserNotifications(filter)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponsePagination(c, constants.GetNotificationsSuccess, pagination, response.ListUserNotificationResponse{}.FromListEntity(result))
}

func (h userNotificationHandler) DeleteNotification(c echo.Context) error {
	id := c.Param("id")

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := h.service.DeleteNotification(id, claims.UserID); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.DeleteNotificationSuccess, nil)
}

func (h userNotificationHandler) MarkNotificationAsRead(c echo.Context) error {
	id := c.Param("id")

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := h.service.MarkNotificationAsRead(id, claims.UserID); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.MarkNotificationAsReadSuccess, nil)
}

func (h userNotificationHandler) MarkAllNotificationsAsRead(c echo.Context) error {
	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := h.service.MarkAllNotificationsAsRead(claims.UserID); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.MarkAllNotificationAsReadSuccess, nil)
}
