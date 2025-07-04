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

func (cu *CreateBloodOxygenation) Execute(esp32ID string, bloodOxygenation float64) error {
	err := cu.db.SaveBloodOxygenation(esp32ID, bloodOxygenation)
	if err != nil {
		return err
	}

	bloodOxygenations := domain.NewBloodOxygenation(esp32ID, bloodOxygenation)

	err = cu.rabbit.Save(bloodOxygenations)
	if err != nil {
		return err
	}
	
	return nil
}