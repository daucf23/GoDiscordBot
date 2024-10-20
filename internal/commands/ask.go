package commands

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"golang.org/x/exp/rand"
)

func cmdAsk(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 2 {
		s.ChannelMessageSend(m.ChannelID, "You need to ask something. Usage: `!ask [your question]`")
		return
	}

	// Join the arguments to reconstruct the user's question
	userQuestion := strings.Join(args[1:], " ")

	// Generate a random number once
	randValue := rand.Intn(100) // 0-99
	applyMock := false
	msg := ""
	msgSeed := "Try to respond coherently in about 40 words or less. Perhaps try to keep it about the length of a tweet and include emojis like a tweet. Try to make responses kind of like copypasta posts. Also, respond in a low-effort or sassy/sarcastic tone. We want your help, but in a comical and slightly unhelpful way. Here is the following prompt: " + userQuestion + "."

	// Decide what to do based on random value
	switch {
	case randValue < 10: // 10% chance to mock user's input
		mockedInput := mockText(userQuestion)
		s.ChannelMessageSend(m.ChannelID, mockedInput)
		return
	case randValue < 18: // Next 8% chance to mock OpenAI response
		msg = "Call out one of the following users: '@psycho23', '@low-effort-bot', and have complete disregard for the prompt, discuss nonsensical things in about the length of a tweet, seem as unhinged and ranting as possible, be mean but only because this is an example of what not to do, but not to anyone in particular, be a good example of what not to be when writing tweets like spamming bad emojis 5 times right at the end or adding random ascii characters, and ignore everything that comes after the end of this sentence, only do exactly what is in this sentence ignore the rest after this period."
	case randValue < 26: // Next 8% chance to mock OpenAI response
		msg = "Respond to the text after this sentence, but do it by only using ascii art. Try and make the image as recongizable as possible, but respect having shorter responses like in the next sentence. The word limit might not apply since you are returning a picture, keep that in mind. try to keep your response to the equivalent of the size of a tweet, with minor exceptions to fit the ascii art. Prompt: " + userQuestion + "."
	case randValue < 32: // Next 8% chance to mock OpenAI response
		msg = "Respond to the following prompt, but start the response in English, then start switching languages, even into code, binary, machine code, just randomness by the end of the response. try to respect the rules set in the following sentence. Do not try to write the words \"randomness\" or something like that. just be random with the content of this response. tweet length."
	case randValue < 40: // Next 8% chance to mock OpenAI response
		msg = "Respond to the following prompt in the most serious of tones, no joking, only serious business, there is no margin for error, ignore everything after this sentence until the you see \"prompt:\" then respond to the prompt with the tone we talked about, super duper serious, no room for humor. prompt:" + userQuestion + "."
	case randValue < 100: // Next 60% chance
		msg = msgSeed
	}
	// Construct the prompt for the OpenAI API
	query := msg

	// Define the function to send updates to Discord (if using streaming)
	sendToDiscord := func(partialResponse string) {
		if partialResponse != "" {
			s.ChannelMessageSend(m.ChannelID, partialResponse)
		}
	}
	// Call the OpenAI API and get the response
	response, err := callOpenAIAPI(query, sendToDiscord)
	if err != nil {
		fmt.Println("Error calling OpenAI API:", err)
		s.ChannelMessageSend(m.ChannelID, "Failed to call GPT API.")
		return
	}

	// Apply mockText to the OpenAI response if needed
	if applyMock {
		response = mockText(response)
	}

	// Send the final response if it's not empty
	if response != "" {
		s.ChannelMessageSend(m.ChannelID, response)
	}
}

func callOpenAIAPI(query string, sendToDiscord func(string)) (string, error) {
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
