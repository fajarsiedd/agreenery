package auth

import (
	"go-agreenery/entities"
	"go-agreenery/middlewares"
)

type AuthService interface {
	Login(user entities.User) (entities.User, error)
	Register(user entities.User) (entities.User, error)
	GetNewTokens(claims *middlewares.JWTCustomClaims, refreshToken string) (entities.User, error)
	GetProfile(id string) (entities.User, error)
	UpdateProfile(user entities.User, selectedFields []string) (entities.User, error)
}
