package infraestructure

import (
	"net/http"
	"vitalPoint/src/body-temperature/application"

	"github.com/gin-gonic/gin"
)

type CreateBodyTemperatureController struct {
	useCase *application.CreateBodyTemperature
}

func NewCreateBodyTemperatureController(useCase *application.CreateBodyTemperature) *CreateBodyTemperatureController {
	return &CreateBodyTemperatureController{useCase: useCase}
}

type RequestBody struct {
	Esp32ID      string  `json:"esp32_id"`
	Tiempo       string  `json:"tiempo"`
	TempAmbiente float64 `json:"temp_ambiente"`
	TempObjeto   float64 `json:"temp_objeto"`
}

func (ct_c *CreateBodyTemperatureController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	// Si no viene el Esp32ID, ponle un valor por defecto
	if body.Esp32ID == "" {
		body.Esp32ID = "1ESP32"
	}

	err := ct_c.useCase.Execute(body.Esp32ID, body.Tiempo, body.TempAmbiente, body.TempObjeto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la temperatura corporal", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Temperatura corporal agregada correctamente"})
}
