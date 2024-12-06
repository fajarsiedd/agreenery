package models

import (
	"go-agreenery/constants/enums"
	"go-agreenery/entities"
)

type Credential struct {
	Base
	Email    string `gorm:"unique"`
	Password string
	Role     enums.Role `gorm:"type:enum('user', 'admin');column:role;default:'user'"`
}

func (c Credential) FromEntity(credentialEntity entities.Credential) Credential {
	return Credential{
		Base:     c.Base.FromEntity(credentialEntity.Base),
		Email:    credentialEntity.Email,
		Password: credentialEntity.Password,
		Role:     credentialEntity.Role,
	}
}

func (c Credential) ToEntity() entities.Credential {
	return entities.Credential{
		Base:     c.Base.ToEntity(),
		Email:    c.Email,
		Password: c.Password,
		Role:     c.Role,
	}
}
