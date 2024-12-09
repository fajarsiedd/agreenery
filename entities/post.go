package entities

import "mime/multipart"

type Post struct {
	Base
	UserID        string
	User          User
	Content       string
	Media         string
	MediaFile     multipart.File
	CategoryID    string
	Category      Category
	CountLikes    int64
	Likes         []Like
	CountComments int64
	Comments      []Comment
	IsLiked       bool
}
