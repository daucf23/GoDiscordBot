package internal

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func CallOpenAIAPI(query string) (string, error) {
	// Get API key from environment variables
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("missing OPENAI_API_KEY environment variable")
	}

	// Initialize OpenAI client
	client := openai.NewClient(option.WithAPIKey(apiKey), option.WithHeader("OpenAI-Beta", "assistants=v2"))

	ctx := context.Background()

	fmt.Println("API Key: ", apiKey)
	fmt.Println("Query: ", query)

	// Existing Assistant ID
	assistantID := "asst_82obYP26qtottx6GKC4yCUPn"

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

	for stream.Next() {
		evt := stream.Current()

		// Debug: print out the received data
		//fmt.Printf("Received data: %+v\n", evt.Data)

		// Handle different event types
		switch eventData := evt.Data.(type) {
		case openai.Run:
			fmt.Println("Run metadata received, status:", eventData.Status)
			// You might want to handle `Run` events differently, but they mostly contain metadata.
		case openai.RunStep:
			fmt.Println("Run step received, step details:", eventData.StepDetails)
		case openai.MessageDeltaEvent:
			// Here is where the response content is streamed
			for _, contentPart := range eventData.Delta.Content {
				if contentPart.Type == "text" {
					response += contentPart.Text.Value // Collect text content from the delta
				}
			}
		case openai.Message:
			// Final complete message (may or may not be used, depending on your use case)
			if eventData.Status == "completed" {
				fmt.Println("Message completed, content:", response)
			}
		default:
			fmt.Printf("Unexpected event data type: %T\n", evt.Data)
		}
	}

	// Check for stream errors after iteration
	if stream.Err() != nil {
		return "", fmt.Errorf("error in stream: %v", stream.Err())
	}

	return response, nil

}
