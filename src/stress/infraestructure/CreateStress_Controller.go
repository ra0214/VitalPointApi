package infraestructure

import (
	"net/http"
	"vitalPoint/src/stress/application"

	"github.com/gin-gonic/gin"
)

type CreateStressController struct {
	useCase *application.CreateStress
}

func NewCreateStressController(useCase *application.CreateStress) *CreateStressController {
	return &CreateStressController{useCase: useCase}
}

type RequestBody struct {
	Esp32ID     string  `json:"esp32_id"`
	Timestamp   string  `json:"tiempo"`
	Temperatura float64 `json:"temperatura"`
	Oxigenacion float64 `json:"oxigenacion"`
}

func (ct_c *CreateStressController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	// Si no viene el Esp32ID, ponle un valor por defecto
	if body.Esp32ID == "" {
		body.Esp32ID = "1ESP32"
	}

	err := ct_c.useCase.Execute(body.Esp32ID, body.Timestamp, body.Temperatura, body.Oxigenacion)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el estres", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Estres agregado correctamente"})
}
