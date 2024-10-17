package commands

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func CallOpenAIAPI(query string, sendToDiscord func(string)) (string, error) {
	// Get API key from environment variables
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("missing OPENAI_API_KEY environment variable")
	}

	// Initialize OpenAI client
	client := openai.NewClient(option.WithAPIKey(apiKey), option.WithHeader("OpenAI-Beta", "assistants=v2"))

	// Increase the timeout to handle long-running tasks (e.g., 15 minutes)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
	defer cancel()

	fmt.Println("API Key: ", apiKey)
	fmt.Println("Query: ", query)

	// Existing Assistant ID
	assistantID := "asst_q5iV50WBjQ1QuRzTp6rQMbkQ"

	// Create a thread
	thread, err := client.Beta.Threads.New(ctx, openai.BetaThreadNewParams{})
	if err != nil {
		return "", fmt.Errorf("error creating thread: %v", err)
	}

	// Create a message in the thread
	_, err = client.Beta.Threads.Messages.New(ctx, thread.ID, openai.BetaThreadMessageNewParams{
		Role: openai.F(openai.BetaThreadMessageNewParamsRoleUser),
		Content: openai.F([]openai.MessageContentPartParamUnion{
			openai.TextContentBlockParam{
				Type: openai.F(openai.TextContentBlockParamTypeText),
				Text: openai.String(query),
			},
		}),
	})
	if err != nil {
		return "", fmt.Errorf("error creating message: %v", err)
	}

	// Use the existing assistant to create a run (streaming response)
	stream := client.Beta.Threads.Runs.NewStreaming(ctx, thread.ID, openai.BetaThreadRunNewParams{
		AssistantID:  openai.String(assistantID),
		Instructions: openai.String("Respond to the user's query."),
	})

	// Check for error after stream initialization
	if stream.Err() != nil {
		return "", fmt.Errorf("error creating run: %v", stream.Err())
	}

	var response string
	var buffer string           // Buffer to accumulate parts of the response
	const maxBufferLength = 500 // Adjust to the number of characters to send per chunk

	for stream.Next() {
		evt := stream.Current()

		// Handle different event types
		switch eventData := evt.Data.(type) {
		case openai.MessageDeltaEvent:
			// Here is where the response content is streamed
			for _, contentPart := range eventData.Delta.Content {
				if contentPart.Type == "text" {
					response += contentPart.Text.Value
					buffer += contentPart.Text.Value

					// Send buffer to Discord if it's reached the chunk size limit
					if len(buffer) >= maxBufferLength && strings.HasSuffix(buffer, ".") {
						sendToDiscord(buffer)
						buffer = "" // Clear the buffer
					}
				}
			}
		case openai.Message:
			if eventData.Status == "completed" {
				fmt.Println("Message completed, content:", response)
			}
		default:
			fmt.Printf("Unexpected event data type: %T\n", evt.Data)
		}
	}

	// Flush any remaining buffer after the stream ends
	if buffer != "" {
		sendToDiscord(buffer)
	}

	// Check for stream errors after iteration
	if stream.Err() != nil {
		return "", fmt.Errorf("error in stream: %v", stream.Err())
	}
	response = ""
	return response, nil
}
