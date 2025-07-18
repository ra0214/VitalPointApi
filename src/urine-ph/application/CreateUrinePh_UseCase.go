package application

import (
	"vitalPoint/src/urine-ph/domain"
)

type CreateUrinePh struct {
	db domain.IUrinePh
	rabbit domain.IUrinePhRabbitMQ
}

func NewCreateUrinePh(db domain.IUrinePh, r domain.IUrinePhRabbitMQ) *CreateUrinePh {
	return &CreateUrinePh{db: db, rabbit: r}
}

func (cu *CreateUrinePh) Execute(esp32ID string, tiempo string, ph int32) error {
	err := cu.db.SaveUrinePh(esp32ID, tiempo, ph)
	if err != nil {
		return err
	}

	urinePh := domain.NewUrinePh(esp32ID, tiempo, ph)

	err = cu.rabbit.Save(urinePh)
	if err != nil {
		return err
	}
	
	return nil
}