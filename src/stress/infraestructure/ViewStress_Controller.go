package infraestructure

import (
	"net/http"
	"vitalPoint/src/stress/application"
	"vitalPoint/src/stress/domain"

	"github.com/gin-gonic/gin"
)

type ViewStressController struct {
	useCase *application.ViewStress
	repo    domain.IStress
}

func NewViewStressController(useCase *application.ViewStress, repo domain.IStress) *ViewStressController {
	return &ViewStressController{
		useCase: useCase,
		repo:    repo,
	}
}

func (vc *ViewStressController) Execute(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")

	// Obtener todos los datos de estrés
	data, err := vc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Error al obtener los datos",
			"detalles": err.Error(),
		})
		return
	}

	// Devolver los datos directamente
	c.JSON(http.StatusOK, data)
}
