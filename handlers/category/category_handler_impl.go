package category

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/request"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"go-agreenery/services/category"

	"github.com/labstack/echo/v4"
)

type categoryHandler struct {
	service category.CategoryService
}

func NewCategoryHandler(s category.CategoryService) *categoryHandler {
	return &categoryHandler{
		service: s,
	}
}

func (h categoryHandler) GetCategories(c echo.Context) error {
	filter, err := helpers.GetFilter(c)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	result, pagination, err := h.service.GetCategories(filter)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponsePagination(c, constants.GetCategoriesSuccess, pagination, response.ListCategoryResponse{}.FromListEntity(result))
}

func (h categoryHandler) GetCategory(c echo.Context) error {
	id := c.Param("id")

	result, err := h.service.GetCategory(id)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.GetCategorySuccess, response.CategoryResponse{}.FromEntity(result))
}

func (h categoryHandler) CreateCategory(c echo.Context) error {
	req := request.CategoryRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	category := req.ToEntity()

	result, err := h.service.CreateCategory(category)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.CreateCategorySuccess, response.CategoryResponse{}.FromEntity(result))
}

func (h categoryHandler) UpdateCategory(c echo.Context) error {
	id := c.Param("id")

	req := request.CategoryRequest{ID: id}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	category := req.ToEntity()

	result, err := h.service.UpdateCategory(category)
	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.UpdateCategorySuccess, response.CategoryResponse{}.FromEntity(result))
}

func (h categoryHandler) DeleteCategory(c echo.Context) error {
	id := c.Param("id")

	if err := h.service.DeleteCategory(id); err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccessResponse(c, constants.DeleteCategorySuccess, nil)
}
