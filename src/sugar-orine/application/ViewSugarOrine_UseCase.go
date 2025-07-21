package application

import (
	"vitalPoint/src/sugar-orine/domain"
)

type ViewSugarOrine struct {
	repo domain.ISugarOrine
}

func NewViewSugarOrine(repo domain.ISugarOrine) *ViewSugarOrine {
	return &ViewSugarOrine{repo: repo}
}

func (vbt *ViewSugarOrine) Execute() ([]domain.SugarOrine, error) {
	return vbt.repo.GetAll()
}