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
	url := "https://chat.egune.com/v1/chat/completions"
	promt := "Your task is to summarize the content of a text inside of <text> tag. Your job is to give short, simple, and accurate summary.\r\n\tThe result should have the following structure: \r\n\t<result>\r\n\t\t<comment>Summery of the user inputs in up to 3 sentences.give response in mongolian<\\/comment>\r\n\t<\\/result>\r\n\tHere is the text:\r\n\t<text>" + ToUnicode(input.Content) + "<\\/text>"
	// sanuulga bichij ogoh summarize hiihdee eniig anhaaraarai geh met promt uusgej bichne
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
		"Authorization": "Bearer tM4GWv8ckExtyuSTHg",
		"Content-Type":  "application/json",
	}
	req.Data = body
	req.Result = &result
	res, err := req.Do()

	return res, result, err
}

func docstoreAPI() {
	apiRequest := oapi.NewRequest("POST", "http://192.168.88.213:8005/parse")
	apiRequest.Headers = map[string]string{"token": "e63fba766214c8f49f375d3184febac94287bfe6"}

	var result []*Document
	apiRequest.Result = &result
	response, err := apiRequest.SendPDF("C:/Users/undra/Downloads/image_2024-11-01_195818518.pdf")
	if err != nil {
		// Handle error
		fmt.Println("Error sending PDF:", err)
	}
	// Process the response as needed
	if response.Data != nil {
		for _, doc := range result {
			if doc != nil { // Check if the pointer is not nil
				fmt.Printf("Title: %s\nContent: %s\n\n", doc.TitleCandidate, doc.File)
			}
		}
	}
}

type Document struct {
	DocType        *string    `json:"doc_type"`
	File           string     `json:"file"`
	GroupType      *string    `json:"group_type"`
	Height         int        `json:"height"`
	IsHTML         bool       `json:"is_html"`
	Log            Log        `json:"log"`
	Number         int        `json:"number"`
	PageTexts      []PageText `json:"page_texts"`
	Signatures     []string   `json:"signatures"`
	Stamps         []string   `json:"stamps"`
	Tables         []Table    `json:"tables"`
	TitleCandidate string     `json:"title_candidate"`
	Titles         []Title    `json:"titles"`
	Width          int        `json:"width"`
}

type Log struct {
	CraftBBoxes     []BBox     `json:"craft_bboxes"`
	FailedBoxes     [][]int    `json:"failed_boxes"`
	HorizontalAngle float64    `json:"horizontal_angle"`
	PDFWords        []string   `json:"pdf_words"`
	SkewAngle       float64    `json:"skew_angle"`
	TextBBoxes      []TextBBox `json:"text_bboxes"`
	TextLines       []TextLine `json:"text_lines"`
	TimeLog         TimeLog    `json:"time_log"`
}

type BBox struct {
	FontHeight int     `json:"font_height"`
	Text       *string `json:"text"`
	X1         int     `json:"x1"`
	X2         int     `json:"x2"`
	Y1         int     `json:"y1"`
	Y2         int     `json:"y2"`
}

type TextBBox struct {
	FontHeight int    `json:"font_height"`
	Text       string `json:"text"`
	X1         int    `json:"x1"`
	X2         int    `json:"x2"`
	Y1         int    `json:"y1"`
	Y2         int    `json:"y2"`
}

type TextLine struct {
	FontHeight int    `json:"font_height"`
	Text       string `json:"text"`
	X1         int    `json:"x1"`
	X2         int    `json:"x2"`
	Y1         int    `json:"y1"`
	Y2         int    `json:"y2"`
}

type TimeLog struct {
	FillingTableCells      float64 `json:"Filling table cells"`
	GroupTextBlocks        float64 `json:"Grouping text blocks"`
	GroupTextLines         float64 `json:"Grouping text lines"`
	LayoutDetection        float64 `json:"Layout detection"`
	OCRExtTextExtraction   float64 `json:"OCR text extraction"`
	OrientationCorrection  float64 `json:"Orientation correction"`
	PDFImageTextExtraction float64 `json:"PDF Image and text extraction"`
	PDFTextExtraction      float64 `json:"PDF text extraction"`
	StudentV3TextDetection float64 `json:"Student v3 text detection"`
	TableDetection         float64 `json:"Table detection"`
	TitleTableFormatTime   float64 `json:"Title and table format time"`
	TotalTime              float64 `json:"Total time"`
	Start                  float64 `json:"start"`
}

type PageText struct {
	FontHeight int    `json:"font_height"`
	Text       string `json:"text"`
	X1         int    `json:"x1"`
	X2         int    `json:"x2"`
	Y1         int    `json:"y1"`
	Y2         int    `json:"y2"`
}

type Table struct {
	DataFrame *string    `json:"_dataframe"`
	HTML      *string    `json:"_html"`
	Box       BBox       `json:"box"`
	ColCount  int        `json:"col_cnt"`
	HasHeader bool       `json:"has_header"`
	HasIndex  bool       `json:"has_index"`
	Merged    bool       `json:"merged"`
	RowCount  int        `json:"row_cnt"`
	Rows      []TableRow `json:"rows"`
	Title     string     `json:"title"`
}

type TableRow struct {
	Cells    []TableCell `json:"cells"`
	IsHeader bool        `json:"is_header"`
}

type TableCell struct {
	ColEndID   int    `json:"col_end_id"`
	ColStartID int    `json:"col_start_id"`
	FontHeight int    `json:"font_height"`
	RowEndID   int    `json:"row_end_id"`
	RowStartID int    `json:"row_start_id"`
	Text       string `json:"text"`
	X1         int    `json:"x1"`
	X2         int    `json:"x 2"`
	Y1         int    `json:"y1"`
	Y2         int    `json:"y2"`
}

type Title struct {
	FontHeight int     `json:"font_height"`
	Text       *string `json:"text"`
	X1         int     `json:"x1"`
	X2         int     `json:"x2"`
	Y1         int     `json:"y1"`
	Y2         int     `json:"y2"`
}
