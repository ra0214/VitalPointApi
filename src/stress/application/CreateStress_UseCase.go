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

// Calcula el nivel de estrés basado en temperatura y oxigenación
func calcularNivelStress(temperatura float64, oxigenacion float64) string {
	if (temperatura > 38.5) || (oxigenacion < 90) {
		return "Alto"
	}
	if (temperatura >= 37.6 && temperatura <= 38.5) || (oxigenacion >= 90 && oxigenacion <= 94) {
		return "Medio"
	}
	if (temperatura >= 36 && temperatura <= 37.5) && (oxigenacion >= 95) {
		return "Bajo"
	}
	return "Medio" // Valor por defecto si no entra en ningún rango exacto
}

// Ahora recibe temperatura y oxigenación, calcula el nivel y lo guarda como string
func (cu *CreateStress) Execute(esp32ID string, tiempo string, temperatura float64, oxigenacion float64) error {
	nivel := calcularNivelStress(temperatura, oxigenacion)

	// Aquí puedes guardar el nivel como string, o mapearlo a un valor si tu dominio lo requiere
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
