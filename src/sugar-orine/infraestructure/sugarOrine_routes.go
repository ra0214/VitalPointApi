package infraestructure

import (
	"vitalPoint/src/sugar-orine/application"
	"vitalPoint/src/sugar-orine/domain"

	"github.com/gin-gonic/gin"
)

func SetupSugarOrineRouter(repo domain.ISugarOrine, rabbitRepo domain.ISugarOrineRabbitMQ) *gin.Engine {
	r := gin.Default()

	CreateSugarOrine := application.NewCreateSugarOrine(repo, rabbitRepo)
	createSugarOrineController := NewCreateSugarOrineController(CreateSugarOrine)

	viewSugarOrine := application.NewViewSugarOrine(repo)
	viewSugarOrineController := NewViewSugarOrineController(viewSugarOrine)

	// Crear instancia del caso de uso y controlador de estadísticas
	getSugarOrineStats := application.NewGetSugarOrineStats(repo)
	getSugarOrineStatsController := NewGetSugarOrineStatsController(getSugarOrineStats)

	r.POST("/sugar", createSugarOrineController.Execute)
	r.GET("/sugar", viewSugarOrineController.Execute)
	r.GET("/sugar/stats", getSugarOrineStatsController.GetStats)

	return r
}
