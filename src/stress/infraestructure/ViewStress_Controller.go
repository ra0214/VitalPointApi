package infraestructure

import (
	"net/http"
	"vitalPoint/src/stress/application"
	"vitalPoint/src/stress/domain"

	"github.com/gin-gonic/gin"
)

type ViewStressController struct {
	useCase *application.ViewStress
	repo    domain.IStress
}

func NewViewStressController(useCase *application.ViewStress, repo domain.IStress) *ViewStressController {
	return &ViewStressController{
		useCase: useCase,
		repo:    repo,
	}
}

func (vc *ViewStressController) Execute(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	esp32ID := c.DefaultQuery("esp32_id", "ESP32_001")

	data, err := vc.repo.GetCorrelationData(esp32ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Error al obtener los datos",
			"detalles": err.Error(),
		})
		return
	}

	// Procesar datos para la visualización
	result := gin.H{
		"correlationData": data,
		"summary": map[string]int{
			"Alto":  0,
			"Medio": 0,
			"Bajo":  0,
		},
	}

	// Contar ocurrencias de cada nivel de estrés
	for _, d := range data {
		result["summary"].(map[string]int)[d.Stress]++
	}

	c.JSON(http.StatusOK, result)
}
