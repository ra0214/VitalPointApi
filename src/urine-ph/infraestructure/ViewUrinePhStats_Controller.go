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
	// Agregar headers CORS
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET")

	readings, err := vc.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Error al obtener datos",
			"detalles": err.Error(),
		})
		return
	}

	if len(readings) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    "No hay suficientes datos para el análisis",
			"detalles": "Se necesitan al menos 3 mediciones",
		})
		return
	}

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
