package response

import (
	"go-agreenery/entities"
	"go-agreenery/dto/base"
)

type RegisterResponse struct {
	base.Base
	DisplayName string `json:"display_name"`
	Phone       string `json:"phone"`
	PhotoUrl    string `json:"photo_url"`
	Email       string `json:"email"`
}

func (registerResponse RegisterResponse) FromEntity(user entities.User) RegisterResponse {
	return RegisterResponse{
		Base:        registerResponse.Base.FromEntity(user.Base),
		DisplayName: user.DisplayName,
		Phone:       user.Phone,
		PhotoUrl:    user.PhotoUrl,
		Email:       user.Credential.Email,
	}
}
