package response

// import (
// 	"go-agreenery/dto/base"
// 	"go-agreenery/entities"
// )

// type LikeResponse struct {
// 	base.Base
// 	UserID string `json:"user_id"`
// 	PostID string `json:"post_id"`
// }

// type ListLikeResponse []LikeResponse

// func (r LikeResponse) FromEntity(like entities.Like) LikeResponse {
// 	return LikeResponse{
// 		Base: 
// 		UserID: like.UserID,
// 		PostID: like.PostID,
// 	}
// }

// func (lr ListLikeResponse) FromListEntity(likes []entities.Like) ListLikeResponse {
// 	data := ListLikeResponse{}

// 	for _, v := range likes {
// 		data = append(data, LikeResponse{}.FromEntity(v))
// 	}

// 	return data
// }
