package entities

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
	CraftBBoxes     []BBox      `json:"craft_bboxes"`
	FailedBoxes     [][]int     `json:"failed_boxes"`
	HorizontalAngle float64     `json:"horizontal_angle"`
	PDFWords        interface{} `json:"pdf_words"` // Should be an array of strings
	SkewAngle       float64     `json:"skew_angle"`
	TextBBoxes      []TextBBox  `json:"text_bboxes"`
	TextLines       []TextLine  `json:"text_lines"`
	TimeLog         TimeLog     `json:"time_log"`
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
