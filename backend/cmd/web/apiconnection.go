package main

import (
	"fmt"

	"undrakh.net/summarizer/pkg/common/oapi"
)

type body struct {
	Model              string  `json:"model"`
	Prompt             string  `json:"prompt"`
	Temperatures       int     `json:"temperatures"`
	Repetition_penalty float32 `json:"repetition_penalty"`
	Stream             string  `json:"stream"`
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

func summarizeAPI(input Content) (*oapi.APIResponse, error) {
	url := "https://test.egune.com/v1/completions"

	body := body{
		Model:              "egune",
		Prompt:             input.Content,
		Temperatures:       0,
		Repetition_penalty: 1.1,
		Stream:             "false",
	}

	req := oapi.NewRequest("POST", url)
	req.Headers = map[string]string{
		"Authorization": "h3FyBvIY694Yo382HqfRumUxPpVS7TERyOCgPg7xK1ERqSWz3gsTwL9zC4ovf2QQhjAK31cXjo2pyMhUXHN53u0R4nZIerOnSgq1kkZ7usrLpugDcU6DtxcekXFT1oRm",
		"Content-Type":  "application/json",
	}
	req.Data = body
	res, err := req.Do()

	return res, err
}
