package application

import (
	"fmt"
	"time"
	"vitalPoint/src/stress/domain"
)

type GetDataAndCalculateStress struct {
	db     domain.IStress
	rabbit domain.IStressRabbitMQ
}

func NewGetDataAndCalculateStress(db domain.IStress, r domain.IStressRabbitMQ) *GetDataAndCalculateStress {
	return &GetDataAndCalculateStress{db: db, rabbit: r}
}

type StressData struct {
	ESP32ID     string  `json:"esp32_id"`
	Temperatura float64 `json:"temperatura"`
	Oxigenacion float64 `json:"oxigenacion"`
	StressLevel string  `json:"stress_level"`
	Timestamp   string  `json:"timestamp"`
}

// Solo obtiene los datos sin guardar
func (gdc *GetDataAndCalculateStress) GetData(esp32ID string) (*StressData, error) {
	// Obtener últimos datos de temperatura y oxigenación
	temperatura, err := gdc.db.GetLatestTemperature(esp32ID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener última temperatura: %v", err)
	}

	oxigenacion, err := gdc.db.GetLatestOxygenation(esp32ID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener última oxigenación: %v", err)
	}

	// Calcular nivel de estrés usando la función compartida
	nivel := CalcularNivelStress(temperatura, oxigenacion)

	// Timestamp actual
	tiempo := time.Now().Format("2006-01-02T15:04:05Z07:00")

	return &StressData{
		ESP32ID:     esp32ID,
		Temperatura: temperatura,
		Oxigenacion: oxigenacion,
		StressLevel: nivel,
		Timestamp:   tiempo,
	}, nil
}

// Guarda los datos de estrés calculados
func (gdc *GetDataAndCalculateStress) SaveStress(data *StressData) error {
	// Guardar en base de datos
	err := gdc.db.SaveStress(data.ESP32ID, data.Timestamp, data.StressLevel)
	if err != nil {
		return fmt.Errorf("error guardando en BD: %v", err)
	}

	// Enviar a RabbitMQ
	stress := domain.NewStress(data.ESP32ID, data.Timestamp, data.StressLevel)
	err = gdc.rabbit.Save(stress)
	if err != nil {
		return fmt.Errorf("error enviando a RabbitMQ: %v", err)
	}

	return nil
}
