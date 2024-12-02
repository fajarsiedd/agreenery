package auth

import "go-agreenery/entities"

type AuthRepository interface {
	Login(user entities.User) (entities.User, error)
	Register(user entities.User) (entities.User, error)
	FindUser(id string) (entities.User, error)
	UpdateUser(user entities.User, selectedFields []string) (entities.User, error)
}
