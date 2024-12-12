package chatbot

import "github.com/labstack/echo/v4"

type ChatbotHandler interface {
	SendPrompt(c echo.Context) error
}