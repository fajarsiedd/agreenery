package entities

import "mime/multipart"

type Plant struct {
	Base
	Name         string
	Description  string
	Image        string
	Fertilizer   string
	PlantingTips string
	ImageFile    multipart.File
	CategoryID   string
	Category     Category
	Steps        []Step
}
