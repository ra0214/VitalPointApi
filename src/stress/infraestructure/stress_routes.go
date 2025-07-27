package infraestructure

import (
	"vitalPoint/src/stress/application"
	"vitalPoint/src/stress/domain"

	"github.com/gin-gonic/gin"
)

func SetupStressRouter(repo domain.IStress, rabbitRepo domain.IStressRabbitMQ) *gin.Engine {
	r := gin.Default()

	// Casos de uso existentes
	CreateStress := application.NewCreateStress(repo, rabbitRepo)
	createStressController := NewCreateStressController(CreateStress)

	viewStress := application.NewViewStress(repo)
	viewStressController := NewViewStressController(viewStress)

	// ❌ DESACTIVADO: Obtener datos y calcular estrés manualmente
	// getDataStress := application.NewGetDataAndCalculateStress(repo, rabbitRepo)
	// getDataStressController := NewGetDataStressController(getDataStress)

	// Rutas existentes
	r.POST("/stress", createStressController.Execute)
	r.GET("/stress", viewStressController.Execute)

	// ❌ DESACTIVADO: rutas para manejo manual
	// r.GET("/stress/data", getDataStressController.GetData)       // Obtener datos
	// r.POST("/stress/save", getDataStressController.SaveStress)   // Guardar estrés

	return r
}
