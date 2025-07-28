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
		esp32ID := c.DefaultQuery("esp32_id", "ESP32_001")

		temp, err := repo.GetLatestTemperature(esp32ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error al obtener temperatura",
				"details": err.Error(),
			})
			return
		}

		oxy, err := repo.GetLatestOxygenation(esp32ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error al obtener oxigenación",
				"details": err.Error(),
			})
			return
		}

		stress, err := repo.GetLatestStress(esp32ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Error al obtener estrés",
				"details": err.Error(),
			})
			return
		}

		// Validar que los datos sean válidos
		if temp <= 0 || oxy <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "Datos inválidos",
				"details": "Temperatura u oxigenación con valores no válidos",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"correlationData": []gin.H{
				{
					"esp32_id":    esp32ID,
					"temperatura": temp,
					"oxigenacion": oxy,
					"stress":      stress.Stress,
					"timestamp":   stress.Timestamp,
				},
			},
		})
	})

	return r
}
