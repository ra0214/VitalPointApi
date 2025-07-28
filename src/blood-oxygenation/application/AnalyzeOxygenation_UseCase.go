package application

import (
	"fmt"
	"sort"
	"vitalPoint/src/blood-oxygenation/domain"
)

type OxygenStats struct {
	Media           float64   `json:"media"`
	Mediana         float64   `json:"mediana"`
	Moda            float64   `json:"moda"`
	DesvEstandar    float64   `json:"desviacion_estandar"`
	Minimo          float64   `json:"minimo"`
	Maximo          float64   `json:"maximo"`
	Intervalos      []string  `json:"intervalos"`
	FrecRelativa    []float64 `json:"frecRelativa"`
	FrecAcumulada   []float64 `json:"frecAcumulada"`
	RangosNiveles   []string  `json:"rangosNiveles"`
	PorcentajeRango []float64 `json:"porcentajeRango"`
}

type AnalyzeOxygenation struct {
	repo domain.IBloodOxygenation
}

func NewAnalyzeOxygenation(repo domain.IBloodOxygenation) *AnalyzeOxygenation {
	return &AnalyzeOxygenation{repo: repo}
}

func (a *AnalyzeOxygenation) Execute() (*OxygenStats, error) {
	readings, err := a.repo.GetAll()
	if err != nil {
		return nil, err
	}

	values := make([]float64, 0)
	for _, r := range readings {
		if r.SpO2.Valid {
			values = append(values, r.SpO2.Float64)
		}
	}

	if len(values) == 0 {
		return nil, fmt.Errorf("no hay datos válidos para analizar")
	}

	sort.Float64s(values)

	stats := &OxygenStats{
		Media:        calcularMedia(values),
		Mediana:      calcularMediana(values),
		Moda:         calcularModa(values),
		DesvEstandar: calcularDesvEstandar(values),
		Minimo:       values[0],
		Maximo:       values[len(values)-1],
	}

	intervalos, frecRelativa, frecAcumulada := calcularFrecuencias(values, 10)
	stats.Intervalos = intervalos
	stats.FrecRelativa = frecRelativa
	stats.FrecAcumulada = frecAcumulada

	stats.RangosNiveles = []string{
		"Normal (95-100%)",
		"Leve (90-94%)",
		"Moderado (85-89%)",
		"Severo (<85%)",
	}
	stats.PorcentajeRango = calcularPorcentajesRango(values)

	return stats, nil
}
