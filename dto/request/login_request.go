package request

import "go-agreenery/entities"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r LoginRequest) ToEntity() entities.User {
	return entities.User{
		Credential: entities.Credential{
			Email:    r.Email,
			Password: r.Password,
		},
	}
}
