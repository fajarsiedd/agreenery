package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type ProfileResponse struct {
	base.Base
	DisplayName  string `json:"display_name"`
	Phone        string `json:"phone"`
	PhotoProfile string `json:"photo_profile"`
	Email        string `json:"email"`
}

func (loginResponse ProfileResponse) FromEntity(user entities.User) ProfileResponse {
	return ProfileResponse{
		Base:         loginResponse.Base.FromEntity(user.Base),
		DisplayName:  user.DisplayName,
		Phone:        user.Phone,
		PhotoProfile: user.PhotoProfile,
		Email:        user.Credential.Email,
	}
}
