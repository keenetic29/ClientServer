package domain

type FileUploadRequest struct {
    FileName string `json:"file_name"`
    Content  string `json:"content"`
}