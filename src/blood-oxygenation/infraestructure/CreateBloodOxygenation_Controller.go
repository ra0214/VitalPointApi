package infraestructure

import (
	"net/http"
	"vitalPoint/src/blood-oxygenation/application"

	"github.com/gin-gonic/gin"
)

type CreateBloodOxygenationController struct {
	useCase *application.CreateBloodOxygenation
}

func NewCreateBloodOxygenationController(useCase *application.CreateBloodOxygenation) *CreateBloodOxygenationController {
	return &CreateBloodOxygenationController{useCase: useCase}
}

type RequestBody struct {
	Esp32ID   string `json:"esp32_id"`
	Timestamp string `json:"tiempo"`
	IR        int32  `json:"ir"`
	Red       int32  `json:"red"`
	SpO2      float32  `json:"spo2"`
}

func (ct_c *CreateBloodOxygenationController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	// Si no viene el Esp32ID, ponle un valor por defecto
	if body.Esp32ID == "" {
		body.Esp32ID = "1ESP32"
	}

	err := ct_c.useCase.Execute(body.Esp32ID, body.Timestamp, body.IR, body.Red, body.SpO2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la saturación de oxígeno en sangre", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Saturación de oxígeno en sangre agregada correctamente"})
}
