package infraestructure

import (
	"net/http"
	"vitalPoint/src/stress/application"
	"vitalPoint/src/stress/domain"

	"github.com/gin-gonic/gin"
)

func SetupStressRouter(repo domain.IStress, rabbitRepo domain.IStressRabbitMQ) *gin.Engine {
	r := gin.Default()

	// Configurar CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Casos de uso existentes
	CreateStress := application.NewCreateStress(repo, rabbitRepo)
	createStressController := NewCreateStressController(CreateStress)

	viewStress := application.NewViewStress(repo)
	viewStressController := NewViewStressController(viewStress, repo)

	// Rutas existentes
	r.POST("/stress", createStressController.Execute)
	r.GET("/stress", viewStressController.Execute)

	// Ruta para datos de correlación
	r.GET("/stress/correlation", func(c *gin.Context) {
		esp32ID := c.DefaultQuery("esp32_id", "1ESP32")

		correlationData, err := repo.GetCorrelationData(esp32ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error al obtener datos de correlación",
				"details": err.Error(),
			})
			return
		}

		// Transformar los datos al formato esperado
		var formattedData []gin.H
		for _, d := range correlationData {
			formattedData = append(formattedData, gin.H{
				"esp32_id":    d.ESP32ID,
				"temperatura": d.Temperatura,
				"oxigenacion": d.Oxigenacion,
				"stress":      d.Stress,
				"timestamp":   d.Timestamp,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"correlationData": formattedData,
		})
	})

	return r
}
