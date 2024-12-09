package models

import "go-agreenery/entities"

type Like struct {
	Base
	UserID string `gorm:"size:191"`
	PostID string `gorm:"size:191"`
}

type ListLike []Like

func (l Like) FromEntity(like entities.Like) Like {
	return Like{
		Base:   l.Base.FromEntity(like.Base),
		UserID: like.UserID,
		PostID: like.PostID,
	}
}

func (l Like) ToEntity() entities.Like {
	return entities.Like{
		Base:   l.Base.ToEntity(),
		UserID: l.UserID,
		PostID: l.PostID,
	}
}

func (ll ListLike) FromListEntity(likes []entities.Like) ListLike {
	data := ListLike{}

	for _, v := range likes {
		data = append(data, Like{}.FromEntity(v))
	}

	return data
}

func (ll ListLike) ToListEntity() []entities.Like {
	data := []entities.Like{}

	for _, v := range ll {
		data = append(data, v.ToEntity())
	}

	return data
}
