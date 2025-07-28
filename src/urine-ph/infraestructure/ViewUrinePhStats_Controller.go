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
	c.Header("Access-Control-Allow-Origin", "*")

	readings, err := vc.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"grupos_horarios": []interface{}{
				map[string]interface{}{
					"periodo":             "Mañana",
					"media":               0,
					"desviacion_estandar": 0,
					"n":                   0,
				},
				map[string]interface{}{
					"periodo":             "Tarde",
					"media":               0,
					"desviacion_estandar": 0,
					"n":                   0,
				},
				map[string]interface{}{
					"periodo":             "Noche",
					"media":               0,
					"desviacion_estandar": 0,
					"n":                   0,
				},
			},
			"estadistico_f":             0,
			"valor_p":                   0,
			"significancia_estadistica": false,
		})
		return
	}

	stats, err := vc.useCase.Execute(readings)
	if err != nil {
		c.JSON(http.StatusOK, stats) // Devolver los stats vacíos en lugar de error
		return
	}

	c.JSON(http.StatusOK, stats)
}
