package auth

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/request"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"go-agreenery/middlewares"
	"go-agreenery/services/auth"
	"io"

	"github.com/h2non/filetype"
	"github.com/labstack/echo/v4"
)

type authHandler struct {
	service auth.AuthService
}

func NewAuthHandler(s auth.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h authHandler) Login(c echo.Context) error {
	req := request.LoginRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	result, err := h.service.Login(req.ToEntity())
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.LoginSuccess, response.LoginResponse{}.FromEntity(result))
}

func (h authHandler) Register(c echo.Context) error {
	req := request.RegisterRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	result, err := h.service.Register(req.ToEntity())
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.RegisterSuccess, response.RegisterResponse{}.FromEntity(result))
}

func (h authHandler) GetNewTokens(c echo.Context) error {
	claims, token, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	result, err := h.service.GetNewTokens(claims, token)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GenerateTokenSuccess, response.RefreshTokenResponse{}.FromEntity(result))
}

func (h authHandler) GetProfile(c echo.Context) error {
	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	result, err := h.service.GetProfile(claims.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetProfileSuccess, response.ProfileResponse{}.FromEntity(result))
}

func (h authHandler) UpdateProfile(c echo.Context) error {
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

	result, err := h.service.UpdateProfile(req.ToEntity(), req.ToCleanFields())
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.UpdateProfileSuccess, response.ProfileResponse{}.FromEntity(result))
}

func (h authHandler) UploadProfilePhoto(c echo.Context) error {
	file, err := c.FormFile("file")
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

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	user, err := h.service.UploadProfilePhoto(blobFile, claims.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.UploadProfilePhotoSuccess, response.ProfileResponse{}.FromEntity(user))
}
