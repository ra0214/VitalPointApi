package infraestructure

import (
	"vitalPoint/src/blood-oxygenation/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewBloodOxygenationController struct {
	useCase *application.ViewBloodOxygenation
}

func NewViewBloodOxygenationController(useCase *application.ViewBloodOxygenation) *ViewBloodOxygenationController {
	return &ViewBloodOxygenationController{useCase: useCase}
}

func (vc *ViewBloodOxygenationController) Execute(c *gin.Context) {
	data, err := vc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
