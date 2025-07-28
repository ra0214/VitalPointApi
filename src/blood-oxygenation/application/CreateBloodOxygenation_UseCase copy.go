package application

import (
	"vitalPoint/src/blood-oxygenation/domain"
)

type CreateBloodOxygenation struct {
	db     domain.IBloodOxygenation
	rabbit domain.IBloodOxygenationRabbitMQ
}

func NewCreateBloodOxygenation(db domain.IBloodOxygenation, r domain.IBloodOxygenationRabbitMQ) *CreateBloodOxygenation {
	return &CreateBloodOxygenation{db: db, rabbit: r}
}

func (cu *CreateBloodOxygenation) Execute(esp32ID string, tiempo string, ir int32, red int32, spo2 float64) error {
	// Ajustar el valor de SpO2 según la regla
	adjustedSpo2 := spo2
	if spo2 < 100 && spo2 <= 90 {
		adjustedSpo2 = spo2 + 10
	}

	// Guardar con el valor ajustado
	err := cu.db.SaveBloodOxygenation(esp32ID, tiempo, ir, red, adjustedSpo2)
	if err != nil {
		return err
	}

	// Crear objeto con el valor ajustado
	bloodOxygenations := domain.NewBloodOxygenation(esp32ID, tiempo, ir, red, adjustedSpo2)

	err = cu.rabbit.Save(bloodOxygenations)
	if err != nil {
		return err
	}

	return nil
}
