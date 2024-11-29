package request

import "go-agreenery/entities"

type RegisterRequest struct {
	DisplayName string `json:"display_name" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
}

func (registerRequest RegisterRequest) ToEntity() entities.User {
	return entities.User{
		DisplayName: registerRequest.DisplayName,
		Phone:       registerRequest.Phone,
		Credential: entities.Credential{
			Email:    registerRequest.Email,
			Password: registerRequest.Password,
		},
	}
}
