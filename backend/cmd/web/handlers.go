package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"undrakh.net/summarizer/cmd/web/app"
	"undrakh.net/summarizer/pkg/common/oapi"
)

type Content struct {
	Content string `json:"content"`
	Summary string `json:"summary"`
}

func ping(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("OK"))
}

func echo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.Copy(w, r.Body)
}

func logout(w http.ResponseWriter, r *http.Request) {
	app.Session.Remove(r, "auth_user_id")
	app.Session.Remove(r, "oauth2_provider_name")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func clearSession(w http.ResponseWriter, r *http.Request) {
	app.Session.Remove(r, "auth_user_id")
	app.Session.Remove(r, "oauth2_provider_name")
	oapi.SendResp(w, "OK")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func pdfConvert(w http.ResponseWriter, r *http.Request) {
	// Check that the Content-Type is application/pdf
	if r.Header.Get("Content-Type") != "application/pdf" {
		oapi.ClientError(w, http.StatusUnsupportedMediaType)
		return
	}

	// Read the raw PDF data from the request body
	var pdfData bytes.Buffer
	if _, err := io.Copy(&pdfData, r.Body); err != nil {
		http.Error(w, "Failed to read PDF data", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// pdfData.Bytes() contains the raw PDF bytes
	fmt.Printf("Received PDF of size %d bytes\n", pdfData.Len())

	// Send a response back to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "PDF received successfully"}`))
}

type FilePathRequest struct {
	Path string `json:"path"`
}

const UploadDirectory = "./uploads"

// uploadPDFHandler handles PDF uploads from the frontend.
func uploadPDFHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the request is of type application/pdf
	if r.Header.Get("Content-Type") != "application/pdf" {
		http.Error(w, "Invalid content type", http.StatusBadRequest)
		return
	}

	// Ensure upload directory exists, create if necessary
	if _, err := os.Stat(UploadDirectory); os.IsNotExist(err) {
		if err := os.MkdirAll(UploadDirectory, os.ModePerm); err != nil {
			http.Error(w, "Failed to create upload directory", http.StatusInternalServerError)
			return
		}
	}

	// Generate a unique filename based on the current timestamp
	filename := fmt.Sprintf("uploaded_%d.pdf", time.Now().Unix())
	filePath := filepath.Join(UploadDirectory, filename)

	// Create a new file to save the uploaded PDF
	file, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Copy the PDF data from the request body to the file
	if _, err := io.Copy(file, r.Body); err != nil {
		http.Error(w, "Failed to save file content", http.StatusInternalServerError)
		return
	}

	// Respond with a success message and file path
	response := fmt.Sprintf("File saved successfully at %s", filePath)
	w.Write([]byte(response))
}
func summarizer(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Unable to read body: %v", err)
		oapi.CustomError(w, http.StatusBadRequest, "Unable to read body")
		return
	}
	defer r.Body.Close()

	var content Content
	err = json.Unmarshal(body, &content)
	if err != nil {
		log.Printf("Invalid JSON: %v", err)
		oapi.CustomError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	resp, chatResp, err := summarizeAPI(content)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}

	if resp.ErrMessage != "" {
		oapi.CustomError(w, resp.Code, resp.ErrMessage)
		return
	}

	if len(chatResp.Choices) == 0 {
		oapi.CustomError(w, http.StatusInternalServerError, "No choices returned from summarization")
		return
	}

	content.Summary = chatResp.Choices[0].Message.Content
	response, err := json.Marshal(content)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
