package domain

type FileAnalysis struct {
	FileName string `json:"file_name"`
	Lines    int    `json:"lines"`
	Words    int    `json:"words"`
	Chars    int    `json:"chars"`
}

type FileUploadRequest struct {
	FileName string `json:"file_name"`
	Content  string `json:"content"`
}