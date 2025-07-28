package infraestructure

import (
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

	// Nueva ruta para datos de correlación
	r.GET("/stress/correlation", func(c *gin.Context) {
		esp32ID := c.DefaultQuery("esp32_id", "ESP32_001")

		// Validar que el ESP32ID no esté vacío
		if esp32ID == "" {
			c.JSON(400, gin.H{
				"error": "ESP32ID no puede estar vacío",
			})
			return
		}

		data, err := repo.GetCorrelationData(esp32ID)
		if err != nil {
			c.JSON(500, gin.H{
				"error":   "Error al obtener datos de correlación",
				"details": err.Error(),
			})
			return
		}

		// Validar que haya datos
		if len(data) == 0 {
			c.JSON(404, gin.H{
				"error": "No se encontraron datos de correlación",
			})
			return
		}

		// Procesar los datos
		result := gin.H{
			"correlationData": data,
			"summary": map[string]int{
				"Alto":  0,
				"Medio": 0,
				"Bajo":  0,
			},
		}

		// Contar ocurrencias de cada nivel
		for _, d := range data {
			if nivel, ok := result["summary"].(map[string]int); ok {
				nivel[d.Stress]++
			}
		}

		c.JSON(200, result)
	})

	return r
}
