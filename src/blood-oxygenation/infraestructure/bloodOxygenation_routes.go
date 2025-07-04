package infraestructure

import (
	"vitalPoint/src/blood-oxygenation/application"
	"vitalPoint/src/blood-oxygenation/domain"

	"github.com/gin-gonic/gin"
)

func SetupBloodOxygenationRouter(repo domain.IBloodOxygenation, rabbitRepo domain.IBloodOxygenationRabbitMQ) *gin.Engine {
	r := gin.Default()

	CreateBloodOxygenation := application.NewCreateBloodOxygenation(repo, rabbitRepo)
	createBloodOxygenationController := NewCreateBloodOxygenationController(CreateBloodOxygenation)

	viewBloodOxygenation := application.NewViewBloodOxygenation(repo)
	viewBloodOxygenationController := NewViewBloodOxygenationController(viewBloodOxygenation)

	r.POST("/bloodOxygenation", createBloodOxygenationController.Execute)
	r.GET("/bloodOxygenation", viewBloodOxygenationController.Execute)

	return r
}
