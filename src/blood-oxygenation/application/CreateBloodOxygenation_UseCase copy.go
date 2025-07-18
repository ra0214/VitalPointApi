package application

import (
	"vitalPoint/src/blood-oxygenation/domain"
)

type CreateBloodOxygenation struct {
	db domain.IBloodOxygenation
	rabbit domain.IBloodOxygenationRabbitMQ
}

func NewCreateBloodOxygenation(db domain.IBloodOxygenation, r domain.IBloodOxygenationRabbitMQ) *CreateBloodOxygenation {
	return &CreateBloodOxygenation{db: db, rabbit: r}
}

func (cu *CreateBloodOxygenation) Execute(esp32ID string, tiempo string, ir int32, red int32, spo2 float32) error {
	err := cu.db.SaveBloodOxygenation(esp32ID, tiempo, ir, red, spo2)
	if err != nil {
		return err
	}

	bloodOxygenations := domain.NewBloodOxygenation(esp32ID, tiempo, ir, red, spo2)

	err = cu.rabbit.Save(bloodOxygenations)
	if err != nil {
		return err
	}
	
	return nil
}