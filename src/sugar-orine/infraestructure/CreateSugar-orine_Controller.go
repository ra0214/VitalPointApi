package infraestructure

import (
	"net/http"
	"vitalPoint/src/sugar-orine/application"

	"github.com/gin-gonic/gin"
)

type CreateSugarOrineController struct {
	useCase *application.CreateSugarOrine
}

func NewCreateSugarOrineController(useCase *application.CreateSugarOrine) *CreateSugarOrineController {
	return &CreateSugarOrineController{useCase: useCase}
}

type RequestBody struct {
	Esp32ID   string  `json:"esp32_id"`
	Timestamp string  `json:"tiempo"`
	Glucosa   string  `json:"glucosa"`
}

func (ct_c *CreateSugarOrineController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	// Si no viene el Esp32ID, ponle un valor por defecto
	if body.Esp32ID == "" {
		body.Esp32ID = "1ESP32"
	}

	err := ct_c.useCase.Execute(body.Esp32ID, body.Timestamp, body.Glucosa)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el nivel de glucosa", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Nivel de glucosa agregado correctamente"})
}
