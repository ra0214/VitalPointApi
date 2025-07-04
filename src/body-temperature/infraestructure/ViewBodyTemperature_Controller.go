package infraestructure

import (
	"vitalPoint/src/body-temperature/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewBodyTemperatureController struct {
	useCase *application.ViewBodyTemperature
}

func NewViewBodyTemperatureController(useCase *application.ViewBodyTemperature) *ViewBodyTemperatureController {
	return &ViewBodyTemperatureController{useCase: useCase}
}

func (vc *ViewBodyTemperatureController) Execute(c *gin.Context) {
	data, err := vc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
