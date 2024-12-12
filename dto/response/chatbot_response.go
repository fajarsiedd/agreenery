package response

import "go-agreenery/entities"

type ChatbotResponse struct {
	Prompt   string `json:"prompt"`
	Response string `json:"response"`
}

func (r ChatbotResponse) FromEntity(chatbot entities.Chatbot) ChatbotResponse {
	return ChatbotResponse{
		Prompt:   chatbot.Prompt,
		Response: chatbot.Response,
	}
}
