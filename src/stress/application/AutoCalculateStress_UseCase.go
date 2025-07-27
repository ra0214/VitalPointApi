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

// Calcula el nivel de estrés basado en temperatura y oxigenación
func calcularNivelStressFromData(temperatura float64, oxigenacion float64) string {
	if (temperatura > 38.5) || (oxigenacion < 90) {
		return "Alto"
	}
	if (temperatura >= 37.6 && temperatura <= 38.5) || (oxigenacion >= 90 && oxigenacion <= 94) {
		return "Medio"
	}
	if (temperatura >= 36 && temperatura <= 37.5) && (oxigenacion >= 95) {
		return "Bajo"
	}
	return "Medio" // Valor por defecto
}

// Calcula el estrés basado en los últimos datos guardados
func (acs *AutoCalculateStress) Execute(esp32ID string) error {
	// Obtener últimos datos de temperatura y oxigenación
	temperatura, err := acs.db.GetLatestTemperature(esp32ID)
	if err != nil {
		return fmt.Errorf("error al obtener última temperatura: %v", err)
	}

	oxigenacion, err := acs.db.GetLatestOxygenation(esp32ID)
	if err != nil {
		return fmt.Errorf("error al obtener última oxigenación: %v", err)
	}

	// Solo calcular si tenemos datos válidos
	if temperatura == 0 && oxigenacion == 0 {
		return fmt.Errorf("no hay datos válidos para calcular estrés")
	}

	// Calcular nivel de estrés
	nivel := calcularNivelStressFromData(temperatura, oxigenacion)

	// Timestamp actual
	tiempo := time.Now().Format("2006-01-02T15:04:05Z07:00")

	// Guardar en base de datos
	err = acs.db.SaveStress(esp32ID, tiempo, nivel)
	if err != nil {
		return err
	}

	// Enviar a RabbitMQ
	stress := domain.NewStress(esp32ID, tiempo, nivel)
	err = acs.rabbit.Save(stress)
	if err != nil {
		return err
	}

	return nil
}

// Ejecutar cálculo automático cada minuto
func (acs *AutoCalculateStress) StartAutoCalculation(esp32ID string) {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		defer ticker.Stop()

		for range ticker.C {
			err := acs.Execute(esp32ID)
			if err != nil {
				fmt.Printf("Error calculando estrés automático: %v\n", err)
			} else {
				fmt.Printf("Estrés calculado automáticamente para ESP32: %s\n", esp32ID)
			}
		}
	}()
}
