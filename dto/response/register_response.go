package response

import (
	"go-agreenery/dto/base"
	"go-agreenery/entities"
)

type RegisterResponse struct {
	base.Base
	DisplayName  string `json:"display_name"`
	Phone        string `json:"phone"`
	PhotoProfile string `json:"photo_profile"`
	Email        string `json:"email"`
}

func (registerResponse RegisterResponse) FromEntity(user entities.User) RegisterResponse {
	return RegisterResponse{
		Base:         registerResponse.Base.FromEntity(user.Base),
		DisplayName:  user.DisplayName,
		Phone:        user.Phone,
		PhotoProfile: user.PhotoProfile,
		Email:        user.Credential.Email,
	}
}
