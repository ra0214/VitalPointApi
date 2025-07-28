package application

import (
	"vitalPoint/src/sugar-orine/domain"
)

type GetSugarOrineStats struct {
	db domain.ISugarOrine
}

func NewGetSugarOrineStats(db domain.ISugarOrine) *GetSugarOrineStats {
	return &GetSugarOrineStats{db: db}
}

func (g *GetSugarOrineStats) Execute() (*domain.SugarOrineStats, error) {
	return g.db.GetStats()
}
