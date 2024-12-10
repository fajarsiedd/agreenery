package models

import (
	"go-agreenery/entities"
)

type User struct {
	Base
	DisplayName  string
	Phone        string `gorm:"unique"`
	Photo        string
	CredentialID string     `gorm:"size:191"`
	Credential   Credential `gorm:"foreignKey:CredentialID;references:ID"`
	Posts        ListPost
	Comments     ListComment
	Likes        ListLike
	Article      ListArticle
	PostReport   ListPostReport
}

func (u User) FromEntity(user entities.User) User {
	return User{
		Base:         u.Base.FromEntity(user.Base),
		DisplayName:  user.DisplayName,
		Phone:        user.Phone,
		Photo:        user.Photo,
		CredentialID: user.CredentialID,
		Credential:   u.Credential.FromEntity(user.Credential),
	}
}

func (u User) ToEntity() entities.User {
	return entities.User{
		Base:         u.Base.ToEntity(),
		DisplayName:  u.DisplayName,
		Phone:        u.Phone,
		Photo:        u.Photo,
		CredentialID: u.CredentialID,
		Credential:   u.Credential.ToEntity(),
	}
}
