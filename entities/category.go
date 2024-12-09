package entities

import "go-agreenery/constants/enums"

type Category struct {
	Base
	Name       string
	Type       enums.CategoryType
	Plants     []Plant
	Article    []Article
	Posts      []Post
	CountPosts int64
}
