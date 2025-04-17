package transport

import (
    "bytes"
    "encoding/json"
    "net/http"
    "client/internal/domain"
)

type HTTPClient struct {
    baseURL string
}

func New(baseURL string) *HTTPClient {
    return &HTTPClient{baseURL: baseURL}
}

func (c *HTTPClient) SendAnalysis(req domain.FileUploadRequest) (*domain.FileAnalysisResponse, error) {
    jsonBody, _ := json.Marshal(req)
    resp, err := http.Post(c.baseURL+"/analyze", "application/json", bytes.NewBuffer(jsonBody))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result domain.FileAnalysisResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    return &result, nil
}