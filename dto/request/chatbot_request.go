package request

import (
	"go-agreenery/entities"
)

type ChatbotRequest struct {
	Prompt      string `form:"prompt" validate:"required"`
	Image       []byte
	ImageFormat string
}

func (r ChatbotRequest) ToEntity() entities.Chatbot {
	return entities.Chatbot{
		Prompt:      r.Prompt,
		Image:       r.Image,
		ImageFormat: r.ImageFormat,
	}
}
