package helpers

import (
	"context"
	"fmt"
	"go-agreenery/entities"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GetFertilzerAndPlantingRecommendation(plant string) (string, string) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	prompt1 := "Short answer only. For development."
	prompt2 := "Tanaman: " + plant
	prompt3 := "Jika tanaman ada berikan rekomendasi jumlah pupuk untuk tanaman dalam bentuk angka dan berikan rekomendasi tentang mananam tanamannya."
	prompt4 := "PENTING: Langsung ke jawaban, jangan gunakan awalan seperti 'Selada:'. pisahkan jawaban menggunakan '#'."
	prompt5 := "PENTING: Jika tanaman tidak ada balas dengan teks persis besar kecil hurufnya 'Jenis tanaman tidak ditemukan' jangan berikan detail jawaban."

	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx,
		genai.Text(prompt1),
		genai.Text(prompt2),
		genai.Text(prompt3),
		genai.Text(prompt4),
		genai.Text(prompt5),
	)

	if err != nil {
		log.Fatal(err)
	}

	result := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				result += fmt.Sprintf("%s", part)
			}
		}
	}

	fertilizer := "Jenis tanaman tidak ditemukan"
	planting := "Jenis tanaman tidak ditemukan"

	if !strings.Contains(result, "Jenis tanaman tidak ditemukan") {
		result = strings.ReplaceAll(result, "\n", "")
		splittedTxt := strings.Split(result, "#")
		fertilizer = strings.TrimSpace(splittedTxt[0])
		planting = strings.TrimSpace(splittedTxt[1])
	}

	return fertilizer, planting
}

func GetResponseFromGemini(chatbot entities.Chatbot) entities.Chatbot {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash")

	var resp *genai.GenerateContentResponse
	if len(chatbot.Image) > 0 && chatbot.ImageFormat != "" {
		resp, err = model.GenerateContent(ctx,
			genai.ImageData(chatbot.ImageFormat, chatbot.Image),
			genai.Text(chatbot.Prompt),
		)
	} else {
		resp, err = model.GenerateContent(ctx,
			genai.Text(chatbot.Prompt),
		)
	}

	if err != nil {
		log.Fatal(err)
	}

	result := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				result += fmt.Sprintf("%s", part)
			}
		}
	}

	return entities.Chatbot{
		Prompt:   chatbot.Prompt,
		Response: result,
	}
}
