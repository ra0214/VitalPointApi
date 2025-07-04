package application

import (
	"vitalPoint/src/body-temperature/domain"
)

type CreateBodyTemperature struct {
	db domain.IBodyTemperature
	rabbit domain.IBodyTemperatureRabbitMQ
}

func NewCreateBodyTemperature(db domain.IBodyTemperature, r domain.IBodyTemperatureRabbitMQ) *CreateBodyTemperature {
	return &CreateBodyTemperature{db: db, rabbit: r}
}

func (cu *CreateBodyTemperature) Execute(esp32ID string, tiempo string, temp_ambiente float64, temp_objeto float64) error {
	err := cu.db.SaveBodyTemperature(esp32ID, tiempo, temp_ambiente, temp_objeto)
	if err != nil {
		return err
	}

	bodyTemperature := domain.NewBodyTemperature(esp32ID, tiempo, temp_ambiente, temp_objeto)

	err = cu.rabbit.Save(bodyTemperature)
	if err != nil {
		return err
	}
	
	return nil
}