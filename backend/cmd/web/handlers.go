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

func uploadPDFHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/pdf" {
		oapi.ClientError(w, http.StatusUnsupportedMediaType)
		return
	}

	if _, err := os.Stat(UploadDirectory); os.IsNotExist(err) {
		if err := os.MkdirAll(UploadDirectory, os.ModePerm); err != nil {
			oapi.ServerError(w, err)
			return
		}
	}

	filename := fmt.Sprintf("uploaded_%d.pdf", time.Now().Unix())
	filePath := filepath.Join(UploadDirectory, filename)

	file, err := os.Create(filePath)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}
	defer file.Close()

	if _, err := io.Copy(file, r.Body); err != nil {
		oapi.ServerError(w, err)
		return
	}
	req := Pdf{
		FilePath: filePath,
	}

	resp, result, err := docstoreAPI(req)

	if err != nil {
		oapi.ServerError(w, err)
		return
	}
	if resp.ErrMessage != "" {
		oapi.CustomError(w, resp.Code, resp.ErrMessage)
		return
	}

	Lecture, err := OCR_result(result)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}

	content := &Content{
		Content: "",
		Summary: "",
	}

	for len(Lecture) > 0 {
		chunkSize := 3500

		// Adjust chunkSize if remaining content is less than chunkSize
		if len(Lecture) < chunkSize {
			chunkSize = len(Lecture)
		}

		// Get the chunk and tokenize it
		chunk := Lecture[:chunkSize]
		tokens := tokenize(chunk) // Replace with actual tokenization method

		// Ensure we don't cut off mid-token
		for i := len(tokens) - 1; i >= 0; i-- {
			if len(tokens[:i+1]) <= chunkSize { // Check if the token count is within limit
				chunk = joinTokens(tokens[:i+1]) // Join back to string
				break
			}
		}

		content.Content = chunk // Set the content for this iteration

		// Call the summarization API
		resp, chatResp, err := summarizeAPI(*content)
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

		// Append the summary
		content.Summary += chatResp.Choices[0].Message.Content

		// Remove the processed chunk from Lecture
		Lecture = Lecture[len(chunk):]
	}

	response, err := json.Marshal(content)
	if err != nil {
		oapi.ServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
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

// func convertTextToPDF(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodOptions {
// 		w.WriteHeader(http.StatusOK)
// 		return
// 	}

// 	textFilePath := "C:/Users/anuji/Video_summarizer/backend/uploads/textFile/summary.txt"
// 	pdf := gofpdf.New("P", "mm", "A4", "")
// 	pdf.AddPage()

// 	log.Println("Adding font...")
// 	pdf.AddUTF8Font("OpenSans", "", "C:/Users/anuji/Video_summarizer/backend/font/OpenSans-Regular.ttf")
// 	pdf.SetFont("OpenSans", "", 12)

// 	log.Println("Reading text file...")
// 	content, err := os.ReadFile(textFilePath)
// 	if err != nil {
// 		log.Printf("Error reading text file: %v", err)
// 		http.Error(w, "Unable to read text file", http.StatusInternalServerError)
// 		return
// 	}
// 	pdf.MultiCell(0, 10, string(content), "", "", false)

// 	pdfFilePath := filepath.Join("C:/Users/anuji/Video_summarizer/backend/uploads/pdf", "output.pdf")

// 	log.Println("Writing PDF file...")
// 	err = pdf.OutputFileAndClose(pdfFilePath)
// 	if err != nil {
// 		log.Printf("Error creating PDF file at %s: %v", pdfFilePath, err)
// 		http.Error(w, "Unable to create PDF file", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/pdf")
// 	w.Header().Set("Content-Disposition", "attachment; filename=output.pdf")
// 	http.ServeFile(w, r, pdfFilePath)
// }
