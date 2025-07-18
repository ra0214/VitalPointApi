package application

import (
	"vitalPoint/src/urine-ph/domain"
)

type ViewUrinePh struct {
	repo domain.IUrinePh
}

func NewViewUrinePh(repo domain.IUrinePh) *ViewUrinePh {
	return &ViewUrinePh{repo: repo}
}

func (vbt *ViewUrinePh) Execute() ([]domain.UrinePh, error) {
	return vbt.repo.GetAll()
}