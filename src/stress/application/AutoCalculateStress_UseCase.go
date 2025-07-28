package application

import (
	"fmt"
	"time"
	"vitalPoint/src/stress/domain"
)

type AutoCalculateStress struct {
	db     domain.IStress
	rabbit domain.IStressRabbitMQ
}

func NewAutoCalculateStress(db domain.IStress, r domain.IStressRabbitMQ) *AutoCalculateStress {
	return &AutoCalculateStress{db: db, rabbit: r}
}

// Calcula el estrés basado en los últimos datos guardados
func (acs *AutoCalculateStress) Execute(esp32ID string) error {
	fmt.Printf("=== INICIANDO CÁLCULO DE ESTRÉS PARA ESP32: %s ===\n", esp32ID)

	// Obtener últimos datos de temperatura y oxigenación
	temperatura, err := acs.db.GetLatestTemperature(esp32ID)
	if err != nil {
		fmt.Printf("ERROR obteniendo temperatura: %v\n", err)
		return fmt.Errorf("error al obtener última temperatura: %v", err)
	}
	fmt.Printf("Temperatura obtenida: %.2f\n", temperatura)

	oxigenacion, err := acs.db.GetLatestOxygenation(esp32ID)
	if err != nil {
		fmt.Printf("ERROR obteniendo oxigenación: %v\n", err)
		return fmt.Errorf("error al obtener última oxigenación: %v", err)
	}
	fmt.Printf("Oxigenación obtenida: %.2f\n", oxigenacion)

	// Solo calcular si tenemos datos válidos
	if temperatura == 0 && oxigenacion == 0 {
		fmt.Printf("No hay datos válidos - Temperatura: %.2f, Oxigenación: %.2f\n", temperatura, oxigenacion)
		return fmt.Errorf("no hay datos válidos para calcular estrés")
	}

	// Calcular nivel de estrés
	nivel := CalcularNivelStress(temperatura, oxigenacion)
	fmt.Printf("Nivel de estrés calculado: %s\n", nivel)

	// Timestamp actual - FORMATO CORREGIDO PARA COINCIDIR CON TU BD
	tiempo := time.Now().Format("2006-01-02T15:04:05.000000")
	fmt.Printf("Timestamp: %s\n", tiempo)

	// Guardar en base de datos
	fmt.Printf("Guardando en base de datos...\n")
	err = acs.db.SaveStress(esp32ID, tiempo, nivel)
	if err != nil {
		fmt.Printf("ERROR guardando en BD: %v\n", err)
		return err
	}
	fmt.Printf("Guardado exitoso en BD\n")

	// Enviar a RabbitMQ
	fmt.Printf("Enviando a RabbitMQ...\n")
	stress := domain.NewStress(esp32ID, tiempo, nivel)
	err = acs.rabbit.Save(stress)
	if err != nil {
		fmt.Printf("ERROR enviando a RabbitMQ: %v\n", err)
		return err
	}
	fmt.Printf("Enviado exitoso a RabbitMQ\n")

	fmt.Printf("=== CÁLCULO COMPLETADO EXITOSAMENTE ===\n")
	return nil
}

// Ejecutar cálculo automático cada minuto
func (acs *AutoCalculateStress) StartAutoCalculation(esp32ID string) {
	fmt.Printf("Iniciando cálculo automático para ESP32: %s cada 1 minuto\n", esp32ID)

	// Ejecutar inmediatamente al iniciar para probar
	fmt.Println("Ejecutando cálculo inicial...")
	err := acs.Execute(esp32ID)
	if err != nil {
		fmt.Printf("Error en cálculo inicial: %v\n", err)
	}

	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		defer ticker.Stop()

		for range ticker.C {
			fmt.Println("=== TICK DEL TIMER - CALCULANDO ESTRÉS ===")
			err := acs.Execute(esp32ID)
			if err != nil {
				fmt.Printf("Error calculando estrés automático: %v\n", err)
			} else {
				fmt.Printf("Estrés calculado automáticamente para ESP32: %s\n", esp32ID)
			}
		}
	}()
}
