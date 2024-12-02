package auth

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/request"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"go-agreenery/middlewares"
	"go-agreenery/services/auth"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	service auth.AuthService
}

func NewAuthHandler(service auth.AuthService) *authHandler {
	return &authHandler{
		service: service,
	}
}

func (handler authHandler) Login(c echo.Context) error {
	req := request.LoginRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	result, err := handler.service.Login(req.ToEntity())
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.LoginSuccess, response.LoginResponse{}.FromEntity(result))
}

func (handler authHandler) Register(c echo.Context) error {
	req := request.RegisterRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	result, err := handler.service.Register(req.ToEntity())
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.RegisterSuccess, response.RegisterResponse{}.FromEntity(result))
}

func (handler authHandler) GetNewTokens(c echo.Context) error {
	claims, token, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	result, err := handler.service.GetNewTokens(claims, token)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GenerateTokenSuccess, response.RefreshTokenResponse{}.FromEntity(result))
}

func (handler authHandler) GetProfile(c echo.Context) error {
	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	result, err := handler.service.GetProfile(claims.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetProfileSuccess, response.ProfileResponse{}.FromEntity(result))
}

func (handler authHandler) UpdateProfile(c echo.Context) error {
	req := request.UpdateProfileRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if req.Email != "" {
		if err := c.Validate(&req); err != nil {
			return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
		}
	}

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}
	req.ID = claims.UserID

	result, err := handler.service.UpdateProfile(req.ToEntity(), req.ToCleanFields())
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.UpdateProfileSuccess, response.ProfileResponse{}.FromEntity(result))
}
