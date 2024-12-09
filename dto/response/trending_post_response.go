package response

import (
	"go-agreenery/entities"
)

type TrendingPostResponse struct {
	Category   string `json:"category"`
	CountPosts int64  `json:"count_posts"`
}

type ListTrendingPostResponse []TrendingPostResponse

func (r TrendingPostResponse) FromEntity(category entities.Category) TrendingPostResponse {
	return TrendingPostResponse{
		Category:   category.Name,
		CountPosts: category.CountPosts,
	}
}

func (lr ListTrendingPostResponse) FromListEntity(categories []entities.Category) ListTrendingPostResponse {
	data := ListTrendingPostResponse{}

	for _, v := range categories {
		data = append(data, TrendingPostResponse{}.FromEntity(v))
	}

	return data
}
