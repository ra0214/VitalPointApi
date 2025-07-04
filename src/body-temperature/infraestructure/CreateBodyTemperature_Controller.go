package infraestructure

import (
	"vitalPoint/src/body-temperature/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBodyTemperatureController struct {
	useCase *application.CreateBodyTemperature
}

func NewCreateBodyTemperatureController(useCase *application.CreateBodyTemperature) *CreateBodyTemperatureController {
	return &CreateBodyTemperatureController{useCase: useCase}
}

type RequestBody struct {
	Esp32ID   string  `json:"esp32_id"`
	Temperatura float64 `json:"temperatura"`
}

func (ct_c *CreateBodyTemperatureController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	err := ct_c.useCase.Execute(body.Esp32ID ,body.Temperatura)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la temperatura corporal", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Temperatura corporal agregada correctamente"})
}