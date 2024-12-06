package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type ProfileResponse struct {
	base.Base
	DisplayName string `json:"display_name"`
	Phone       string `json:"phone"`
	Photo       string `json:"photo"`
	Email       string `json:"email"`
}

func (r ProfileResponse) FromEntity(user entities.User) ProfileResponse {
	return ProfileResponse{
		Base:        r.Base.FromEntity(user.Base),
		DisplayName: user.DisplayName,
		Phone:       user.Phone,
		Photo:       user.Photo,
		Email:       user.Credential.Email,
	}
}
