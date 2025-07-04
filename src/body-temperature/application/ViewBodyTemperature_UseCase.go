package application

import (
	"vitalPoint/src/body-temperature/domain"
)

type ViewBodyTemperature struct {
	repo domain.IBodyTemperature
}

func NewViewBodyTemperature(repo domain.IBodyTemperature) *ViewBodyTemperature {
	return &ViewBodyTemperature{repo: repo}
}

func (vbt *ViewBodyTemperature) Execute() ([]domain.BodyTemperature, error) {
	return vbt.repo.GetAll()
}