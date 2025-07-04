package infraestructure

import (
	"vitalPoint/src/blood-oxygenation/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateBloodOxygenationController struct {
	useCase *application.CreateBloodOxygenation
}

func NewCreateBloodOxygenationController(useCase *application.CreateBloodOxygenation) *CreateBloodOxygenationController {
	return &CreateBloodOxygenationController{useCase: useCase}
}

type RequestBody struct {
	Esp32ID       string  `json:"esp32_id"`
	BloodOxygenation float64 `json:"blood_oxygenation"`
	Timestamp     string  `json:"tiempo"`
	IR            int32   `json:"ir"`
	Red           int32   `json:"red"`
}

func (ct_c *CreateBloodOxygenationController) Execute(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el JSON", "detalles": err.Error()})
		return
	}

	err := ct_c.useCase.Execute(body.Esp32ID ,body.BloodOxygenation, body.Timestamp, body.IR, body.Red)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al guardar la saturación de oxígeno en sangre", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Saturación de oxígeno en sangre agregada correctamente"})
}