package application

import (
	"vitalPoint/src/stress/domain"
)

type CreateStress struct {
	db     domain.IStress
	rabbit domain.IStressRabbitMQ
}

func NewCreateStress(db domain.IStress, r domain.IStressRabbitMQ) *CreateStress {
	return &CreateStress{db: db, rabbit: r}
}

// Ahora recibe temperatura y oxigenación, calcula el nivel y lo guarda como string
func (cu *CreateStress) Execute(esp32ID string, tiempo string, temperatura float64, oxigenacion float64) error {
	nivel := CalcularNivelStress(temperatura, oxigenacion) // Usar la función compartida

	// Guardar el nivel como string
	err := cu.db.SaveStress(esp32ID, tiempo, nivel)
	if err != nil {
		return err
	}

	stres := domain.NewStress(esp32ID, tiempo, nivel)

	err = cu.rabbit.Save(stres)
	if err != nil {
		return err
	}
	return nil
}
