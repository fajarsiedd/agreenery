package models

import "go-agreenery/entities"

type Comment struct {
	Base
	UserID  string `gorm:"size:191"`
	User    User   `gorm:"foreignKey:UserID;references:ID"`
	PostID  string `gorm:"size:191"`
	Message string
}

type ListComment []Comment

func (c Comment) FromEntity(comment entities.Comment) Comment {
	return Comment{
		Base:    c.Base.FromEntity(comment.Base),
		UserID:  comment.UserID,
		User:    c.User.FromEntity(comment.User),
		PostID:  comment.PostID,
		Message: comment.Message,
	}
}

func (c Comment) ToEntity() entities.Comment {
	return entities.Comment{
		Base:    c.Base.ToEntity(),
		UserID:  c.UserID,
		User:    c.User.ToEntity(),
		PostID:  c.PostID,
		Message: c.Message,
	}
}

func (lc ListComment) FromListEntity(comments []entities.Comment) ListComment {
	data := ListComment{}

	for _, v := range comments {
		data = append(data, Comment{}.FromEntity(v))
	}

	return data
}

func (lc ListComment) ToListEntity() []entities.Comment {
	data := []entities.Comment{}

	for _, v := range lc {
		data = append(data, v.ToEntity())
	}

	return data
}
