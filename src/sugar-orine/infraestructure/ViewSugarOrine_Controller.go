package infraestructure

import (
	"vitalPoint/src/sugar-orine/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewSugarOrineController struct {
	useCase *application.ViewSugarOrine
}

func NewViewSugarOrineController(useCase *application.ViewSugarOrine) *ViewSugarOrineController {
	return &ViewSugarOrineController{useCase: useCase}
}

func (vc *ViewSugarOrineController) Execute(c *gin.Context) {
	data, err := vc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los datos", "detalles": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
