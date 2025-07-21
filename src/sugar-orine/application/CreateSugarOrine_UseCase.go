package application

import (
	"vitalPoint/src/sugar-orine/domain"
)

type CreateSugarOrine struct {
	db domain.ISugarOrine
	rabbit domain.ISugarOrineRabbitMQ
}

func NewCreateSugarOrine(db domain.ISugarOrine, r domain.ISugarOrineRabbitMQ) *CreateSugarOrine {
	return &CreateSugarOrine{db: db, rabbit: r}
}

func (cu *CreateSugarOrine) Execute(esp32ID string, tiempo string, glucosa string) error {
	err := cu.db.SaveSugarOrine(esp32ID, tiempo, glucosa)
	if err != nil {
		return err
	}

	sugarOrine := domain.NewSugarOrine(esp32ID, tiempo, glucosa)

	err = cu.rabbit.Save(sugarOrine)
	if err != nil {
		return err
	}
	
	return nil
}