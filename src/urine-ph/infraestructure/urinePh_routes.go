package infraestructure

import (
	"vitalPoint/src/urine-ph/application"
	"vitalPoint/src/urine-ph/domain"

	"github.com/gin-gonic/gin"
)

func SetupUrinePhRouter(repo domain.IUrinePh, rabbitRepo domain.IUrinePhRabbitMQ) *gin.Engine {
	r := gin.Default()

	CreateUrinePh := application.NewCreateUrinePh(repo, rabbitRepo)
	createUrinePhController := NewCreateUrinePhController(CreateUrinePh)

	viewUrinePh := application.NewViewUrinePh(repo)
	viewUrinePhController := NewViewUrinePhController(viewUrinePh)

	r.POST("/urinepH", createUrinePhController.Execute)
	r.GET("/urinepH", viewUrinePhController.Execute)

	return r
}
