package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const claudeURL = "https://api.anthropic.com/v1/messages"
const openaiURL = "https://api.openai.com/v1/chat/completions"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model     string    `json:"model"`
	Messages  []Message `json:"messages"`
	MaxTokens int64     `json:"max_tokens"`
}

type Response struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a message as a command line argument")
	}

	var userMessage string
	var llm string
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-llm=") {
			llmOption := strings.TrimPrefix(arg, "-llm=")
			if llmOption == "claude" || llmOption == "gpt" {
				llm = llmOption
			}
		} else {
			userMessage += arg + " "
		}
	}
	// LLM specific switching
	os_arg := "ANTHROPIC_API_KEY"
	model := "claude-3-5-sonnet-20240620"
	url := claudeURL
	if llm == "gpt" {
		os_arg = "OPENAI_API_KEY"
		model = "gpt-4o-mini"
		url = openaiURL
	}
	apiKey := os.Getenv(os_arg)
	if apiKey == "" {
		fmt.Printf("Please set the %s environment variable\n", os_arg)
		return
	}

	request := Request{
		Model: model,
		Messages: []Message{
			{Role: "user", Content: userMessage},
		},
		MaxTokens: 2048,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("anthropic-version", "2023-06-01")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}
	if len(response.Content) > 0 && strings.TrimSpace(response.Content[0].Text) != "" {
		fmt.Println(response.Content[0].Text)
	} else if len(response.Choices) > 0 {
		fmt.Println(response.Choices[0].Message.Content)
	} else {
		fmt.Print(response)
		fmt.Println("No response content received")
	}
}
