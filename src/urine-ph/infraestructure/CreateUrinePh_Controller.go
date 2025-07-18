package infraestructure

import (
	"net/http"
	"vitalPoint/src/urine-ph/application"

	"github.com/gin-gonic/gin"
)

type CreateUrinePhController struct {
	useCase *application.CreateUrinePh
}

func NewCreateUrinePhController(useCase *application.CreateUrinePh) *CreateUrinePhController {
	return &CreateUrinePhController{useCase: useCase}
}

type RequestBody struct {
	Esp32ID   string `json:"esp32_id"`
	Timestamp string `json:"tiempo"`
	PH        int32  `json:"ph"`
}

func (ct_c *CreateUrinePhController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	// Si no viene el Esp32ID, ponle un valor por defecto
	if body.Esp32ID == "" {
		body.Esp32ID = "1ESP32"
	}

	err := ct_c.useCase.Execute(body.Esp32ID, body.Timestamp, float64(body.PH))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar el pH de la orina", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "pH de la orina agregado correctamente"})
}
