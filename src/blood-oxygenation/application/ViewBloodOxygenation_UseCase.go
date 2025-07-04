package application

import (
	"vitalPoint/src/blood-oxygenation/domain"
)

type ViewBloodOxygenation struct {
	repo domain.IBloodOxygenation
}

func NewViewBloodOxygenation(repo domain.IBloodOxygenation) *ViewBloodOxygenation {
	return &ViewBloodOxygenation{repo: repo}
}

func (vbt *ViewBloodOxygenation) Execute() ([]domain.BloodOxygenation, error) {
	return vbt.repo.GetAll()
}