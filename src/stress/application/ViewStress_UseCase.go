package application

import (
	"vitalPoint/src/stress/domain"
)

type ViewStress struct {
	repo domain.IStress
}

func NewViewStress(repo domain.IStress) *ViewStress {
	return &ViewStress{repo: repo}
}

func (vbt *ViewStress) Execute() ([]domain.Stress, error) {
	return vbt.repo.GetAll()
}