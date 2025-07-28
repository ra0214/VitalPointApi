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

		if len(correlationData) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "No se encontraron datos de correlación",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"correlationData": correlationData,
		})
	})

	return r
}
