package entities

import "mime/multipart"

type Article struct {
	Base
	Thumbnail     string
	ThumbnailFile multipart.File
	Title         string
	Content       string
	UserID        string
	User          User
	CategoryID    string
	Category      Category
	PublishStatus bool
}
