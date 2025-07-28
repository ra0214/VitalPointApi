package infraestructure

import (
	"net/http"
	"vitalPoint/src/urine-ph/application"
	"vitalPoint/src/urine-ph/domain"

	"github.com/gin-gonic/gin"
)

type ViewUrinePhStatsController struct {
	repo    domain.IUrinePh
	useCase *application.AnalyzeUrinePh
}

func NewViewUrinePhStatsController(repo domain.IUrinePh) *ViewUrinePhStatsController {
	return &ViewUrinePhStatsController{
		repo:    repo,
		useCase: application.NewAnalyzeUrinePh(),
	}
}

func (vc *ViewUrinePhStatsController) Execute(c *gin.Context) {
	// Primero obtenemos los datos del repositorio
	readings, err := vc.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Error al obtener datos",
			"detalles": err.Error(),
		})
		return
	}

	// Luego ejecutamos el análisis con los datos
	stats, err := vc.useCase.Execute(readings)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Error al calcular estadísticas",
			"detalles": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, stats)
}
