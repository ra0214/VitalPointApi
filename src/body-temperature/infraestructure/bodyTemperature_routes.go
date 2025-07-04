package infraestructure

import (
	"vitalPoint/src/body-temperature/application"
	"vitalPoint/src/body-temperature/domain"

	"github.com/gin-gonic/gin"
)

func SetupRouter(repo domain.IBodyTemperature, rabbitRepo domain.IBodyTemperatureRabbitMQ) *gin.Engine {
	r := gin.Default()

	CreateBodyTemperature := application.NewCreateBodyTemperature(repo, rabbitRepo)
	createBodyTemperatureController := NewCreateBodyTemperatureController(CreateBodyTemperature)

	viewBodyTemperature := application.NewViewBodyTemperature(repo)
	viewBodyTemperatureController := NewViewBodyTemperatureController(viewBodyTemperature)

	r.POST("/bodyTemperature/", createBodyTemperatureController.Execute)
	r.GET("/bodyTemperature", viewBodyTemperatureController.Execute)

	return r
}
