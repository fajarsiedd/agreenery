package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type LoginResponse struct {
	base.Base
	DisplayName  string `json:"display_name"`
	Phone        string `json:"phone"`
	Photo        string `json:"photo"`
	Email        string `json:"email"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (loginResponse LoginResponse) FromEntity(user entities.User) LoginResponse {
	return LoginResponse{
		Base:         loginResponse.Base.FromEntity(user.Base),
		DisplayName:  user.DisplayName,
		Phone:        user.Phone,
		Photo:        user.Photo,
		Email:        user.Credential.Email,
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	}
}
