package response

import (
	"go-agreenery/constants/enums"
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type CategoryResponse struct {
	base.Base
	Name string             `json:"name"`
	Type enums.CategoryType `json:"type"`
}

type ListCategoryResponse []CategoryResponse

func (c CategoryResponse) FromEntity(category entities.Category) CategoryResponse {
	return CategoryResponse{
		Base: c.Base.FromEntity(category.Base),
		Name: category.Name,
		Type: category.Type,
	}
}

func (lc ListCategoryResponse) FromListEntity(categories []entities.Category) ListCategoryResponse {
	listCategoryRes := ListCategoryResponse{}

	for _, v := range categories {
		listCategoryRes = append(listCategoryRes, CategoryResponse{}.FromEntity(v))
	}

	return listCategoryRes
}
