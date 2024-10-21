package main

import (
	"fmt"

	"undrakh.net/summarizer/pkg/common/oapi"
)

type body struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"` // Corrected to a slice of Message
}

type Message struct {
	Role      string        `json:"role"`
	Content   string        `json:"content"`
	ToolCalls []interface{} `json:"tool_calls"` // Changed to a slice
}

type Choice struct {
	Index        int         `json:"index"`
	Message      Message     `json:"message"`
	Logprobs     interface{} `json:"logprobs"` // Keep this as is to handle null
	FinishReason string      `json:"finish_reason"`
	StopReason   *string     `json:"stop_reason"` // Keep this as is for handling null
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	TotalTokens      int `json:"total_tokens"`
	CompletionTokens int `json:"completion_tokens"`
}

type ChatResponse struct {
	ID             string      `json:"id"`
	Object         string      `json:"object"`
	Created        int64       `json:"created"`
	Model          string      `json:"model"`
	Choices        []Choice    `json:"choices"`
	Usage          Usage       `json:"usage"`
	PromptLogprobs interface{} `json:"prompt_logprobs"` // Keep this as is to handle null
}

func ToUnicode(message string) string {
	unicodeStr := ""
	for _, r := range message {
		if r == ' ' {
			unicodeStr += " " // Keep space as it is
		} else {
			unicodeStr += fmt.Sprintf("\\u%04x", r) // Convert other characters to Unicode escape sequence
		}
	}
	return unicodeStr
}

func summarizeAPI(input Content) (*oapi.APIResponse, *ChatResponse, error) {
	url := "https://dev-protocol.chimege.com/v1/chat/completions"
	promt := "Your task is to summarize the content of a text inside of <text> tag. Your job is to give short, simple, and accurate summary.\r\n\tThe result should have the following structure: \r\n\t<result>\r\n\t\t<comment>Summery of the user inputs in up to 3 sentences.give response in mongolian<\\/comment>\r\n\t<\\/result>\r\n\tHere is the text:\r\n\t<text>" + ToUnicode(input.Content) + "<\\/text>"

	message := Message{
		Role:    "user",
		Content: promt,
	}
	body := body{
		Model:    "egune",
		Messages: []Message{message},
	}

	var result *ChatResponse
	req := oapi.NewRequest("POST", url)
	req.Headers = map[string]string{
		"Authorization": "Bearer fe6bf18e668b5bb5aff8a3b41eb17f22ec1add6b",
		"Content-Type":  "application/json",
	}
	req.Data = body
	req.Result = &result
	res, err := req.Do()

	return res, result, err
}
