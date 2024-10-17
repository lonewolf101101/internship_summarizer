package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type output struct {
	OutputSentences int    `json:"output_sentences"`
	Providers       string `json:"providers"`
	Text            string `json:"text"`
	Language        string `json:"language"`
}

func (app *application) summarize(input Content) Content {
	url := "https://api.edenai.run/v2/text/summarize"

	payload := output{
		OutputSentences: 3,
		Providers:       "microsoft,connexun,openai,emvista",
		Text:            "Barack Hussein Obama is an American politician who served as the 44th president of the United States from 2009 to 2017.",
		Language:        "en",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMjdiYTYwMzQtNjU1My00NzIzLTg4ZDktZDFiZmNhMWJhYjhlIiwidHlwZSI6ImFwaV90b2tlbiJ9.FmW56SiIAtVf4ZMUNR7mOu8Bxu6nLMvICkO0ysSmhIg")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(res.Body)
		fmt.Printf("Error: %d %s\n", res.StatusCode, string(bodyBytes))
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	fmt.Println("Response:", string(bodyBytes))

	return input
}
