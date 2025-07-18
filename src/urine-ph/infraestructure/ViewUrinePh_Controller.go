package infraestructure

import (
	"vitalPoint/src/urine-ph/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewUrinePhController struct {
	useCase *application.ViewUrinePh
}

func NewViewUrinePhController(useCase *application.ViewUrinePh) *ViewUrinePhController {
	return &ViewUrinePhController{useCase: useCase}
}

func (vc *ViewUrinePhController) Execute(c *gin.Context) {
	data, err := vc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
