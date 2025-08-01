package main

import (
	"log"
	mlxInfra "vitalPoint/src/blood-oxygenation/infraestructure"
	maxInfra "vitalPoint/src/body-temperature/infraestructure"
	"vitalPoint/src/config"
	stressApp "vitalPoint/src/stress/application"
	stress "vitalPoint/src/stress/infraestructure"
	sugar "vitalPoint/src/sugar-orine/infraestructure"
	phInfra "vitalPoint/src/urine-ph/infraestructure"
	userInfra "vitalPoint/src/users/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Inicializar repositorios MYSQL
	userRepo := userInfra.NewMySQL()
	maxRepo := maxInfra.NewMySQL()
	mlxRepo := mlxInfra.NewMySQL()
	phRepo := phInfra.NewMySQL()
	stressRepo := stress.NewMySQL()
	sugarRepo := sugar.NewMySQL()

	rabbitMQRepo, err := config.GetChannel()
	if err != nil {
		log.Fatalf("Error al conectar con RabbitMQ: %v", err)
	}
	defer rabbitMQRepo.Close()

	maxRabbit := maxInfra.NewRabbitRepository(rabbitMQRepo.Ch)
	mlxRabbit := mlxInfra.NewRabbitRepository(rabbitMQRepo.Ch)
	phRabbit := phInfra.NewRabbitRepository(rabbitMQRepo.Ch)
	stressRabbit := stress.NewRabbitRepository(rabbitMQRepo.Ch)
	sugarRabbit := sugar.NewRabbitRepository(rabbitMQRepo.Ch)

	// ✅ ACTIVAR: Cálculo automático de estrés cada minuto
	autoCalculateStress := stressApp.NewAutoCalculateStress(stressRepo, stressRabbit)
	autoCalculateStress.StartAutoCalculation("1ESP32") // Usa el mismo esp32_id que tienes en las tablas
	log.Println("✅ Cálculo automático de estrés iniciado (cada minuto)")

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

	phRouter := phInfra.SetupUrinePhRouter(phRepo, phRabbit)
	for _, route := range phRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	stressRouter := stress.SetupStressRouter(stressRepo, stressRabbit)
	for _, route := range stressRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	sugarRouter := sugar.SetupSugarOrineRouter(sugarRepo, sugarRabbit)
	for _, route := range sugarRouter.Routes() {
		r.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	// Configurar servidor
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Iniciar servidor
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
