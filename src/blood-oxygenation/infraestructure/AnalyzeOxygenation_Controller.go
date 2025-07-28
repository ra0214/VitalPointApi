package infraestructure

import (
	"net/http"
	"vitalPoint/src/blood-oxygenation/application"

	"github.com/gin-gonic/gin"
)

type AnalyzeOxygenationController struct {
	useCase *application.AnalyzeOxygenation
}

func NewAnalyzeOxygenationController(useCase *application.AnalyzeOxygenation) *AnalyzeOxygenationController {
	return &AnalyzeOxygenationController{useCase: useCase}
}

func (ac *AnalyzeOxygenationController) Execute(c *gin.Context) {
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
