package chatbot

import (
	"go-agreenery/constants"
	"go-agreenery/dto/base"
	"go-agreenery/dto/request"
	"go-agreenery/dto/response"
	"go-agreenery/helpers"
	"io"

	"github.com/h2non/filetype"
	"github.com/labstack/echo/v4"
)

type chatbotHandler struct{}

func NewChatbotHandler() *chatbotHandler {
	return &chatbotHandler{}
}

func (h chatbotHandler) SendPrompt(c echo.Context) error {
	req := request.ChatbotRequest{}

	if err := c.Bind(&req); err != nil {
		return base.ErrorResponse(c, err)
	}

	if err := c.Validate(&req); err != nil {
		return base.ErrorResponse(c, helpers.TranslateValidationErr(err))
	}

	file, err := c.FormFile("image")
	if err == nil {
		var maxFileSize int64 = 1048576 * 2
		if file.Size > maxFileSize {
			return base.ErrorResponse(c, constants.ErrFileSizeExceedsLimit)
		}

		temp, _ := file.Open()
		buf, _ := io.ReadAll(temp)
		if !filetype.IsImage(buf) {
			return base.ErrorResponse(c, constants.ErrOnlyImageAllowed)
		}

		kind, _ := filetype.Match(buf)
		if kind == filetype.Unknown {
			return base.ErrorResponse(c, constants.ErrUnkownFileType)
		}

		req.Image = buf
		req.ImageFormat = kind.MIME.Subtype
	}

	chatbot := req.ToEntity()
	resp := helpers.GetResponseFromGemini(chatbot)

	return base.SuccessResponse(c, constants.SendPromptSuccess, response.ChatbotResponse{}.FromEntity(resp))
}
