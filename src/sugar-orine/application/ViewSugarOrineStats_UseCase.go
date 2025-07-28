package application

import (
	"vitalPoint/src/sugar-orine/domain"
)

type ViewSugarOrineStats struct {
	repo domain.ISugarOrine
}

func NewViewSugarOrineStats(repo domain.ISugarOrine) *ViewSugarOrineStats {
	return &ViewSugarOrineStats{repo: repo}
}

func (v *ViewSugarOrineStats) Execute() (*domain.SugarOrineStats, error) {
	return v.repo.GetStats()
}
