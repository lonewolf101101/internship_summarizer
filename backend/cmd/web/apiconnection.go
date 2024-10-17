package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type body struct {
	Model              string `json:"model"`
	Prompt             string `json:"prompt"`
	Matokes            string `json:"matokes"`
	Temperatures       string `json:"temperatures"`
	Repetition_penalty string `json:"repetition_penalty"`
	Stream             string `json:"stream"`
}

func summarize(input Content) Content {
	url := "https://test.egune.com/v1/completions"

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	fmt.Println(string(jsonData))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Set("Authorization", "Bearer h3FyBvIY694Yo382HqfRumUxPpVS7TERyOCgPg7xK1ERqSWz3gsTwL9zC4ovf2QQhjAK31cXjo2pyMhUXHN53u0R4nZIerOnSgq1kkZ7usrLpugDcU6DtxcekXFT1oRm")
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
