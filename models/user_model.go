package models

import (
	"go-agreenery/entities"
)

type User struct {
	Base
	DisplayName  string
	Phone        string `gorm:"unique"`
	PhotoUrl     string
	CredentialID string     `gorm:"size:191"`
	Credential   Credential `gorm:"foreignKey:CredentialID;references:ID"`
}

func (user User) FromEntity(userEntity entities.User) User {
	return User{
		Base:         user.Base.FromEntity(userEntity.Base),
		DisplayName:  userEntity.DisplayName,
		Phone:        userEntity.Phone,
		PhotoUrl:     userEntity.PhotoUrl,
		CredentialID: userEntity.CredentialID,
		Credential:   user.Credential.FromEntity(userEntity.Credential),
	}
}

func (user User) ToEntity() entities.User {
	return entities.User{
		Base:         user.Base.ToEntity(),
		DisplayName:  user.DisplayName,
		Phone:        user.Phone,
		PhotoUrl:     user.PhotoUrl,
		CredentialID: user.CredentialID,
		Credential:   user.Credential.ToEntity(),
	}
}
