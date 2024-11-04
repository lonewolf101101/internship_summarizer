package main

import (

	// "io"
	// "net/http"

	"strings"

	"undrakh.net/summarizer/cmd/web/app"
	"undrakh.net/summarizer/cmd/web/entities"
	"undrakh.net/summarizer/pkg/common/oapi"
)

type Pdf struct {
	FilePath string `json:"file_path"`
	Content  string `json:"content"`
}

type body struct {
	Model    string             `json:"model"`
	Messages []entities.Message `json:"messages"` // Corrected to a slice of Message
}

type Content struct {
	Content string `json:"content"`
	Summary string `json:"summary"`
}

func summarizeAPI(input Content) (*oapi.APIResponse, *entities.ChatResponse, error) {
	url := "https://chat.egune.com/v1/chat/completions"
	basePrompt := "Your task is to summarize the content of the text enclosed within the <text> tags. Please provide a short, simple, and accurate summary in Mongolian, limited to a maximum of 3 sentences.\n\nHere is the text:\n\n<text>"
	// sanuulga bichij ogoh summarize hiihdee eniig anhaaraarai geh met promt uusgej bichne
	var messages []entities.Message
	content := input.Content
	for len(content) > 0 {
		chunkSize := 4000
		if len(content) < chunkSize {
			chunkSize = len(content)
		}

		chunk := content[:chunkSize]
		content = content[chunkSize:]
		if len(messages) == 0 {
			prompt := basePrompt + chunk
			message := entities.Message{
				Role:    "user",
				Content: prompt,
			}
			messages = append(messages, message)
		} else if len(content) == 0 {
			// Last chunk closes the <text> tag
			prompt := chunk + "</text>"
			message := entities.Message{
				Role:    "user",
				Content: prompt,
			}
			messages = append(messages, message)
		} else {
			// Subsequent chunks only include the chunk
			message := entities.Message{
				Role:    "user",
				Content: chunk,
			}
			messages = append(messages, message)
		}
	}
	body := body{
		Model:    "egune",
		Messages: messages,
	}

	var result *entities.ChatResponse
	req := oapi.NewRequest("POST", url)
	req.Headers = map[string]string{
		"Authorization": "Bearer " + app.Config.Summarize_ApiKey,
		"Content-Type":  "application/json",
	}
	req.Data = body
	req.Result = &result
	res, err := req.Do()

	return res, result, err
}

func summarizeLectureAPI(input Content) (*oapi.APIResponse, *entities.ChatResponse, error) {
	url := "https://chat.egune.com/v1/chat/completions"
	promt := "Your task is to analyze the content of the lecture and identify its subsections, even if they are not explicitly marked. Please summarize each subsection by providing clear and concise summaries in Mongolian, limited to a maximum of 3 sentences for each identified subsection within the text enclosed in the <text> tags.\n\nHere is the text:\n\n<text>" + input.Content + "</text>"
	// sanuulga bichij ogoh summarize hiihdee eniig anhaaraarai geh met promt uusgej bichne
	message := entities.Message{
		Role:    "user",
		Content: promt,
	}

	req := oapi.NewRequest("POST", url)
	body := body{
		Model:    "egune",
		Messages: []entities.Message{message},
	}

	var result *entities.ChatResponse
	req.Headers = map[string]string{
		"Authorization": "Bearer " + app.Config.Summarize_ApiKey,
		"Content-Type":  "application/json",
	}
	req.Data = body
	req.Result = &result
	res, err := req.Do()

	return res, result, err
}

func docstoreAPI(input Pdf) (*oapi.APIResponse, []*entities.Document, error) {
	apiRequest := oapi.NewRequest("POST", "http://192.168.88.213:8005/parse")
	apiRequest.Headers = map[string]string{"token": app.Config.OCR_ApiKey}

	var result []*entities.Document
	apiRequest.Result = &result
	response, err := apiRequest.SendPDF(input.FilePath)
	if err != nil {
		return nil, nil, err
	}

	return response, result, err
}
func OCR_result(input []*entities.Document) (string, error) {
	// Initialize the result string
	var result string

	// Iterate over the documents
	for _, document := range input {
		for _, pageText := range document.PageTexts {
			result += pageText.Text + " "
		}
	}

	return result, nil
}

func tokenize(input string) []string {
	// Split the input by spaces or other token delimiters
	return strings.Fields(input) // This is just an example; use an appropriate tokenizer
}

// Your function to join tokens back into a string
func joinTokens(tokens []string) string {
	return strings.Join(tokens, " ") // Join tokens back into a single string
}
