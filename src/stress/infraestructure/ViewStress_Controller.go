package infraestructure

import (
	"vitalPoint/src/stress/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewStressController struct {
	useCase *application.ViewStress
}

func NewViewStressController(useCase *application.ViewStress) *ViewStressController {
	return &ViewStressController{useCase: useCase}
}

func (vc *ViewStressController) Execute(c *gin.Context) {
	data, err := vc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
