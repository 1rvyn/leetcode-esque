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
	"strings"
	"time"
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

	fmt.Println(hintRequest.Code)
	fmt.Println(hintRequest.Language)
	fmt.Println(hintRequest.QuestionID)
	fmt.Println(hintRequest.TestResults)

	openaiClient := openai.NewClient(OpenAi)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		MaxTokens:   200,
		Temperature: 0.1,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleSystem,
				Content: `You are a computer science tutor who is to act as if you are receiving an issue from a student 
				who is attempting to solve a coding question. You must NOT give the student the answer to the question, 
				but instead provide hints to help them fix their code.
				The student's code' has failed test 1 & 3 but has passed test 2.
				The tests are as follows:
				"test_cases": [
					{
					  "input": {
						"nums": [2, 7, 11, 15],
						"target": 9
					  },
					  "output": [0, 1]
					},
					{
					  "input": {
						"nums": [3, 2, 4],
						"target": 6
					  },
					  "output": [1, 2]
					},
					{
					  "input": {
						"nums": [3, 3],
						"target": 6
					  },
					  "output": [0, 1]
					}
				  ]
				Can you help direct the student to fix their code without giving any code examples?`,
			},
		},
		Stream: true,
	}
	stream, err := openaiClient.CreateChatCompletionStream(ctx, req)
	if err != nil {
		log.Printf("ChatCompletionStream error: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Cannot create chat completion stream", "data": err})
	}
	defer stream.Close()

	c.Response().Header.Set("Access-Control-Expose-Headers", "Content-Type")
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

		// Split the response into words
		words := strings.Split(response.Choices[0].Delta.Content, " ")

		// Send each word separately as an SSE event
		for _, word := range words {
			_, _ = c.Write([]byte(fmt.Sprintf("data: %s\n\n", word)))
			// speed of the stream
			time.Sleep(10 * time.Millisecond)
		}
	}

	return nil
}
