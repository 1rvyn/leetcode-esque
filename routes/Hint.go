package routes

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"os"
)

var OpenAi = os.Getenv("OPENAI_API_KEY")

type HintRequest struct {
	Code        string `json:"code"`
	Language    string `json:"language"`
	QuestionID  string `json:"questionID"`
	TestResults string `json:"testResults"`
}

func Hint(c *fiber.Ctx) error {
	fmt.Println("hints endpoint hit")
	var hintRequest HintRequest
	if err := json.Unmarshal(c.Body(), &hintRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Cannot parse JSON", "data": err})
	}

	openaiClient := openai.NewClient(OpenAi)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		MaxTokens:   200,
		Temperature: 0.1,
		Messages:    []openai.ChatCompletionMessage{
			// ... The rest of the request as you've written above
		},
		Stream: true,
	}
	stream, err := openaiClient.CreateChatCompletionStream(ctx, req)
	if err != nil {
		log.Printf("ChatCompletionStream error: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Cannot create chat completion stream", "data": err})
	}
	defer stream.Close()

	c.Response().Header.Set("Content-Type", "text/event-stream")
	c.Response().Header.Set("Cache-Control", "no-cache")
	c.Response().Header.Set("Connection", "keep-alive")

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			log.Printf("Stream error: %v\n", err)
			return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error in chat completion stream", "data": err})
		}

		_, _ = c.Write([]byte(fmt.Sprintf("data: %s\n\n", response.Choices[0].Delta.Content)))
	}

	return nil
}
