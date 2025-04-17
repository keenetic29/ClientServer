package domain

type FileAnalysisResponse struct {
    FileName string `json:"file_name"`
    Lines    int    `json:"lines"`
    Words    int    `json:"words"`
    Chars    int    `json:"chars"`
}