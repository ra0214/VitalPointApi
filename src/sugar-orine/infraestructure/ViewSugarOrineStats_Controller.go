package infraestructure

import (
	"net/http"
	"vitalPoint/src/sugar-orine/application"

	"github.com/gin-gonic/gin"
)

type ViewSugarOrineStatsController struct {
	useCase *application.ViewSugarOrineStats
}

func NewViewSugarOrineStatsController(useCase *application.ViewSugarOrineStats) *ViewSugarOrineStatsController {
	return &ViewSugarOrineStatsController{useCase: useCase}
}

func (vc *ViewSugarOrineStatsController) Execute(c *gin.Context) {
	stats, err := vc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}
