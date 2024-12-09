package post

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/request"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"go-agreenery/middlewares"
	"go-agreenery/services/post"
	"io"

	"github.com/h2non/filetype"
	"github.com/labstack/echo/v4"
)

type postHandler struct {
	service post.PostService
}

func NewPostHandler(s post.PostService) *postHandler {
	return &postHandler{
		service: s,
	}
}

func (h postHandler) GetPosts(c echo.Context) error {
	filter, err := helpers.GetFilter(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	filter.UserID = claims.UserID

	result, pagination, err := h.service.GetPosts(filter)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponsePagination(c, constants.GetPostsSuccess, pagination, response.ListPostResponse{}.FromListEntity(result))
}

func (h postHandler) GetPost(c echo.Context) error {
	id := c.Param("id")

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	result, err := h.service.GetPost(id, claims.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetPostSuccess, response.PostResponse{}.FromEntity(result))
}

func (h postHandler) CreatePost(c echo.Context) error {
	req := request.PostRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	file, err := c.FormFile("media")
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

		req.Media = blobFile
	}

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	req.UserID = claims.UserID
	post := req.ToEntity()

	result, err := h.service.CreatePost(post)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.CreatePostSuccess, response.PostResponse{}.FromEntity(result))
}

func (h postHandler) UpdatePost(c echo.Context) error {
	id := c.Param("id")

	req := request.PostRequest{ID: id}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	file, err := c.FormFile("media")
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

		req.Media = blobFile
	}

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	post := req.ToEntity()

	result, err := h.service.UpdatePost(post, claims.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.UpdatePostSuccess, response.PostResponse{}.FromEntity(result))
}

func (h postHandler) DeletePost(c echo.Context) error {
	id := c.Param("id")

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := h.service.DeletePost(id, claims.UserID); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.DeletePostSuccess, nil)
}

func (h postHandler) LikePost(c echo.Context) error {
	id := c.Param("id")

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	result, err := h.service.LikePost(id, claims.UserID)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.ChangeLikeStatusSuccess, response.PostResponse{}.FromEntity(result))
}

func (h postHandler) GetPostsCountByCategory(c echo.Context) error {
	result, err := h.service.GetPostsCountByCategory()
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.TrendingPostSuccess, response.ListTrendingPostResponse{}.FromListEntity(result))
}
