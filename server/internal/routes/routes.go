package routes

import (
	"server/internal/handlers"
	"server/internal/repository"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Инициализация зависимостей
	repo := repository.NewFileRepository("storage")
	analyzerService := service.NewAnalyzerService(repo)
	analyzeController := handlers.NewAnalyzeController(analyzerService)

	// Маршруты
	r.POST("/analyze", analyzeController.Analyze)

	return r
}