package main

import (
	"github.com/gin-gonic/gin"

	"vitalPoint/src/config"
	"log"
	mlxInfra "vitalPoint/src/blood-oxygenation/infraestructure"
	maxInfra "vitalPoint/src/body-temperature/infraestructure"
	userInfra "vitalPoint/src/users/infraestructure"
)

func main() {
	r := gin.Default()

	// Inicializar repositorios MYSQL
	userRepo := userInfra.NewMySQL()
	maxRepo := maxInfra.NewMySQL()
	mlxRepo := mlxInfra.NewMySQL()

	rabbitMQRepo, err := config.GetChannel()
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}
	defer rabbitMQRepo.Close()

	maxRabbit := maxInfra.NewRabbitRepository(rabbitMQRepo.Ch)
	mlxRabbit := mlxInfra.NewRabbitRepository(rabbitMQRepo.Ch)

	userRouter := userInfra.SetupRouter(userRepo)
	for _, route := range userRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	maxRouter := maxInfra.SetupRouter(maxRepo, maxRabbit)
	for _, route := range maxRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	mlxRouter := mlxInfra.SetupBloodOxygenationRouter(mlxRepo, mlxRabbit)
	for _, route := range mlxRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	// Configurar servidor
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Iniciar servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
