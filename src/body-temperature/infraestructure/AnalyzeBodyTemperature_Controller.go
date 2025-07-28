package infraestructure

import (
	"net/http"
	"vitalPoint/src/body-temperature/application"

	"github.com/gin-gonic/gin"
)

type AnalyzeBodyTemperatureController struct {
	useCase *application.AnalyzeBodyTemperature
}

func NewAnalyzeBodyTemperatureController(useCase *application.AnalyzeBodyTemperature) *AnalyzeBodyTemperatureController {
	return &AnalyzeBodyTemperatureController{useCase: useCase}
}

func (ac *AnalyzeBodyTemperatureController) Execute(c *gin.Context) {
	stats, err := ac.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Error al analizar los datos",
			"detalles": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, stats)
}
