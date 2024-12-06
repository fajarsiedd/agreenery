package request

import "go-agreenery/entities"

type RegisterRequest struct {
	DisplayName string `json:"display_name" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
}

func (r RegisterRequest) ToEntity() entities.User {
	return entities.User{
		DisplayName: r.DisplayName,
		Phone:       r.Phone,
		Credential: entities.Credential{
			Email:    r.Email,
			Password: r.Password,
		},
	}
}
