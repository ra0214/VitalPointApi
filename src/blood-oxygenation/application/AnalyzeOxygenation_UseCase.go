package application

import (
	"fmt"
	"math"
	"sort"
	"vitalPoint/src/blood-oxygenation/domain"
)

type OxygenStats struct {
	Media        float64 `json:"media"`
	Mediana      float64 `json:"mediana"`
	Moda         float64 `json:"moda"`
	DesvEstandar float64 `json:"desviacion_estandar"`
	Minimo       float64 `json:"minimo"`
	Maximo       float64 `json:"maximo"`
	// Datos para la ojiva
	ClasesIntervalos []string  `json:"clasesIntervalos"`
	Frecuencias      []int     `json:"frecuencias"`
	FrecAcumuladas   []int     `json:"frecAcumuladas"`
	PorcentajeAcum   []float64 `json:"porcentajeAcum"`
	NivelesSpO2      []struct {
		Nivel      string  `json:"nivel"`
		Porcentaje float64 `json:"porcentaje"`
	} `json:"nivelesSpO2"`
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

	intervalos, frecuencias, frecAcumuladas, porcentajeAcum := calcularIntervalosOjiva(values)
	stats.ClasesIntervalos = intervalos
	stats.Frecuencias = frecuencias
	stats.FrecAcumuladas = frecAcumuladas
	stats.PorcentajeAcum = porcentajeAcum
	stats.NivelesSpO2 = calcularPorcentajesNiveles(values)

	return stats, nil
}

func calcularIntervalosOjiva(values []float64) ([]string, []int, []int, []float64) {
	// Calcular el número de clases (Regla de Sturges)
	n := len(values)
	k := int(1 + 3.322*math.Log10(float64(n)))

	// Calcular amplitud de clase
	amplitud := (values[len(values)-1] - values[0]) / float64(k)

	clasesIntervalos := make([]string, k)
	frecuencias := make([]int, k)
	frecAcumuladas := make([]int, k)
	porcentajeAcum := make([]float64, k)

	// Calcular frecuencias
	for i := 0; i < k; i++ {
		limInf := values[0] + float64(i)*amplitud
		limSup := limInf + amplitud

		clasesIntervalos[i] = fmt.Sprintf("%.1f-%.1f", limInf, limSup)

		// Contar frecuencias
		for _, v := range values {
			if v >= limInf && (v < limSup || (i == k-1 && v <= limSup)) {
				frecuencias[i]++
			}
		}

		// Calcular frecuencia acumulada
		if i == 0 {
			frecAcumuladas[i] = frecuencias[i]
		} else {
			frecAcumuladas[i] = frecAcumuladas[i-1] + frecuencias[i]
		}

		// Calcular porcentaje acumulado
		porcentajeAcum[i] = (float64(frecAcumuladas[i]) / float64(n)) * 100
	}

	return clasesIntervalos, frecuencias, frecAcumuladas, porcentajeAcum
}

func calcularPorcentajesNiveles(values []float64) []struct {
	Nivel      string  `json:"nivel"`
	Porcentaje float64 `json:"porcentaje"`
} {
	total := float64(len(values))
	niveles := make(map[string]int)

	// Contar valores en cada nivel
	for _, v := range values {
		switch {
		case v >= 95:
			niveles["Normal (95-100%)"]++
		case v >= 90:
			niveles["Leve (90-94%)"]++
		case v >= 85:
			niveles["Moderado (85-89%)"]++
		default:
			niveles["Severo (<85%)"]++
		}
	}

	// Convertir a porcentajes y usar la estructura con tags json
	result := []struct {
		Nivel      string  `json:"nivel"`
		Porcentaje float64 `json:"porcentaje"`
	}{
		{Nivel: "Normal (95-100%)", Porcentaje: float64(niveles["Normal (95-100%)"]) / total * 100},
		{Nivel: "Leve (90-94%)", Porcentaje: float64(niveles["Leve (90-94%)"]) / total * 100},
		{Nivel: "Moderado (85-89%)", Porcentaje: float64(niveles["Moderado (85-89%)"]) / total * 100},
		{Nivel: "Severo (<85%)", Porcentaje: float64(niveles["Severo (<85%)"]) / total * 100},
	}

	return result
}
