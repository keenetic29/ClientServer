package handlers

import (
	"net/http"

	"server/internal/domain"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

type AnalyzeController struct {
	service service.AnalyzerService
}

func NewAnalyzeController(service service.AnalyzerService) *AnalyzeController {
	return &AnalyzeController{service: service}
} 

func (c *AnalyzeController) Analyze(ctx *gin.Context) {
	var req domain.FileUploadRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.service.AnalyzeFile(req.FileName, req.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}