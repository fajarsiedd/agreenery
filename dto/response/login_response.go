package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type LoginResponse struct {
	base.Base
	DisplayName  string `json:"display_name"`
	Phone        string `json:"phone"`
	PhotoProfile string `json:"photo_profile"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (loginResponse LoginResponse) FromEntity(user entities.User) LoginResponse {
	return LoginResponse{
		Base:         loginResponse.Base.FromEntity(user.Base),
		DisplayName:  user.DisplayName,
		Phone:        user.Phone,
		PhotoProfile: user.PhotoProfile,
		Email:        user.Credential.Email,
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	}
}
