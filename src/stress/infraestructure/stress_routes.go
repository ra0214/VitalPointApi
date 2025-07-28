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

	// Crear casos de uso
	calculateStress := application.NewAutoCalculateStress(repo, rabbitRepo)
	getDataStress := application.NewGetDataAndCalculateStress(repo, rabbitRepo)

	// Rutas
	r.GET("/stress/:esp32id", func(c *gin.Context) {
		esp32ID := c.Param("esp32id")
		data, err := getDataStress.GetData(esp32ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, data)
	})

	// Agregar nueva ruta para datos de correlación
	r.GET("/stress/correlation/:esp32id", func(c *gin.Context) {
		esp32ID := c.Param("esp32id")
		data, err := repo.GetCorrelationData(esp32ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	// Iniciar cálculo automático
	go calculateStress.StartAutoCalculation("ESP32_001")

	return r
}
