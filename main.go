package main

import (
	"github.com/gin-gonic/gin"

	userInfra "vitalPoint/src/users/infraestructure"
	"log"
)

func main() {
	r := gin.Default()

	// Inicializar repositorios MYSQL
	userRepo := userInfra.NewMySQL()

	userRouter := userInfra.SetupRouter(userRepo)
	for _, route := range userRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	// Configurar servidor
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Iniciar servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
