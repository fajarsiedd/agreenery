package entities

import "go-agreenery/constants/enums"

type Credential struct {
	Base
	Email        string
	Password     string
	Role         enums.Role
}
