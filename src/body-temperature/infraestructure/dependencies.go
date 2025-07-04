package infraestructure

import (
	"log"
	"vitalPoint/src/config"
)

func Init() {
	mysqlRepo := NewMySQL()

	rabbitMQRepo, err := config.GetChannel()
	if err != nil {
		log.Fatalf("Error al inicializar RabbitMQ: %v", err)
	}
	defer rabbitMQRepo.Close()

	rabbitRepo := NewRabbitRepository(rabbitMQRepo.Ch)

	router := SetupRouter(mysqlRepo, rabbitRepo)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
