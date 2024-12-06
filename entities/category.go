package entities

import "go-agreenery/constants/enums"

type Category struct {
	Base
	Name string
	Type enums.CategoryType
}
