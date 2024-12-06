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

func (r CategoryResponse) FromEntity(category entities.Category) CategoryResponse {
	return CategoryResponse{
		Base: r.Base.FromEntity(category.Base),
		Name: category.Name,
		Type: category.Type,
	}
}

func (lr ListCategoryResponse) FromListEntity(categories []entities.Category) ListCategoryResponse {
	data := ListCategoryResponse{}

	for _, v := range categories {
		data = append(data, CategoryResponse{}.FromEntity(v))
	}

	return data
}
