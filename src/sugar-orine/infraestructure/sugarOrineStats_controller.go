package infraestructure

import (
	"net/http"
	"vitalPoint/src/sugar-orine/application"

	"github.com/gin-gonic/gin"
)

type GetSugarOrineStatsController struct {
	useCase *application.GetSugarOrineStats
}

func NewGetSugarOrineStatsController(useCase *application.GetSugarOrineStats) *GetSugarOrineStatsController {
	return &GetSugarOrineStatsController{useCase: useCase}
}

func (gc *GetSugarOrineStatsController) GetStats(c *gin.Context) {
	stats, err := gc.useCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, stats)
}
