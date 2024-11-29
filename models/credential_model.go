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

func (credential Credential) FromEntity(credentialEntity entities.Credential) Credential {
	return Credential{
		Base:     credential.Base.FromEntity(credentialEntity.Base),
		Email:    credentialEntity.Email,
		Password: credentialEntity.Password,
		Role:     credentialEntity.Role,
	}
}

func (credential Credential) ToEntity() entities.Credential {
	return entities.Credential{
		Base:     credential.Base.ToEntity(),
		Email:    credential.Email,
		Password: credential.Password,
		Role:     credential.Role,
	}
}
