package entities

type Comment struct {
	Base
	UserID  string
	User    User
	PostID  string
	Message string
}
