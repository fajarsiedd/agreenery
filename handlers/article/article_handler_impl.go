package article

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/request"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"go-agreenery/middlewares"
	"go-agreenery/services/article"
	"io"

	"github.com/h2non/filetype"
	"github.com/labstack/echo/v4"
)

type articleHandler struct {
	service article.ArticleService
}

func NewArticleHandler(s article.ArticleService) *articleHandler {
	return &articleHandler{
		service: s,
	}
}

func (h articleHandler) GetArticles(c echo.Context) error {
	filter, err := helpers.GetFilter(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	result, pagination, err := h.service.GetArticles(filter)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponsePagination(c, constants.GetArticlesSuccess, pagination, response.ListArticleResponse{}.FromListEntity(result))
}

func (h articleHandler) GetArticle(c echo.Context) error {
	id := c.Param("id")

	result, err := h.service.GetArticle(id)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetArticleSuccess, response.ArticleResponse{}.FromEntity(result))
}

func (h articleHandler) CreateArticle(c echo.Context) error {
	req := request.ArticleRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	file, err := c.FormFile("thumbnail")
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

		req.Thumbnail = blobFile
	}

	claims, _, err := middlewares.GetCurrentToken(c)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	article := req.ToEntity()
	article.UserID = claims.UserID

	result, err := h.service.CreateArticle(article)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.CreateArticleSuccess, response.ArticleResponse{}.FromEntity(result))
}

func (h articleHandler) UpdateArticle(c echo.Context) error {
	id := c.Param("id")

	req := request.ArticleRequest{ID: id}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	file, err := c.FormFile("thumbnail")
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

		req.Thumbnail = blobFile
	}

	article := req.ToEntity()

	result, err := h.service.UpdateArticle(article)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.UpdateArticleSuccess, response.ArticleResponse{}.FromEntity(result))
}

func (h articleHandler) DeleteArticle(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteArticle(id); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.DeleteArticleSuccess, nil)
}
