package infraestructure

import (
	"net/http"
	"vitalPoint/src/stress/application"

	"github.com/gin-gonic/gin"
)

type GetDataStressController struct {
	useCase *application.GetDataAndCalculateStress
}

func NewGetDataStressController(useCase *application.GetDataAndCalculateStress) *GetDataStressController {
	return &GetDataStressController{useCase: useCase}
}

// Obtener datos de temperatura y oxigenación con estrés calculado
func (gsc *GetDataStressController) GetData(c *gin.Context) {
	esp32ID := c.Query("esp32_id")
	if esp32ID == "" {
		esp32ID = "1ESP32" // Valor por defecto
	}

	data, err := gsc.useCase.GetData(esp32ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Error al obtener datos",
			"detalles": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

// Guardar el nivel de estrés calculado
func (gsc *GetDataStressController) SaveStress(c *gin.Context) {
	var data application.StressData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "Error al leer el JSON",
			"detalles": err.Error(),
		})
		return
	}

	err := gsc.useCase.SaveStress(&data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Error al guardar estrés",
			"detalles": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Estrés guardado correctamente",
		"data":    data,
	})
}
