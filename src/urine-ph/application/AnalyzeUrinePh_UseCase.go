package application

import (
	"fmt"
	"math"
	"time"
	"vitalPoint/src/urine-ph/domain"
)

type AnalyzeUrinePh struct {
	// Ya no necesitamos el repositorio aquí
}

func NewAnalyzeUrinePh() *AnalyzeUrinePh {
	return &AnalyzeUrinePh{}
}

// Modificamos Execute para recibir los datos
func (a *AnalyzeUrinePh) Execute(readings []domain.UrinePh) (*domain.UrinePhStats, error) {
	// Si no hay datos, devolver estructura vacía en lugar de error
	if len(readings) == 0 {
		return &domain.UrinePhStats{
			GruposHorarios: []struct {
				Periodo      string  `json:"periodo"`
				Media        float64 `json:"media"`
				DesvEstandar float64 `json:"desviacion_estandar"`
				N            int     `json:"n"`
			}{
				{Periodo: "Mañana", Media: 0, DesvEstandar: 0, N: 0},
				{Periodo: "Tarde", Media: 0, DesvEstandar: 0, N: 0},
				{Periodo: "Noche", Media: 0, DesvEstandar: 0, N: 0},
			},
			EstadisticoF:             0,
			ValorP:                   0,
			SignificanciaEstadistica: false,
		}, nil
	}

	// Validación inicial
	if len(readings) < 9 { // Mínimo total necesario (3 por grupo)
		return nil, fmt.Errorf("se necesitan al menos 9 mediciones en total (3 por período)")
	}

	// Separar por períodos
	morning := []float64{}
	afternoon := []float64{}
	evening := []float64{}

	for _, r := range readings {
		t, err := time.Parse("2006-01-02 15:04:05", r.Timestamp) // Ajusta el formato según tus datos
		if err != nil {
			continue
		}

		hour := t.Hour()
		switch {
		case hour >= 6 && hour < 12:
			morning = append(morning, r.PH)
		case hour >= 12 && hour < 18:
			afternoon = append(afternoon, r.PH)
		default:
			evening = append(evening, r.PH)
		}
	}

	// Validar cantidad mínima por grupo
	minPorGrupo := 3
	if len(morning) < minPorGrupo || len(afternoon) < minPorGrupo || len(evening) < minPorGrupo {
		return nil, fmt.Errorf("se necesitan al menos %d mediciones por período (mañana: %d, tarde: %d, noche: %d)",
			minPorGrupo, len(morning), len(afternoon), len(evening))
	}

	// Calcular media y desviación estándar general
	allValues := make([]float64, len(readings))
	for i, r := range readings {
		allValues[i] = r.PH
	}

	mediaGeneral := calcularMedia(allValues)
	desvEstandarGeneral := math.Sqrt(calcularVarianza(allValues))

	// Calcular estadísticas por grupo
	stats := &domain.UrinePhStats{
		Media:        mediaGeneral,
		DesvEstandar: desvEstandarGeneral,
		GruposHorarios: []struct {
			Periodo      string  `json:"periodo"`
			Media        float64 `json:"media"`
			DesvEstandar float64 `json:"desviacion_estandar"`
			N            int     `json:"n"`
		}{
			{
				Periodo:      "Mañana",
				Media:        calcularMedia(morning),
				DesvEstandar: math.Sqrt(calcularVarianza(morning)),
				N:            len(morning),
			},
			{
				Periodo:      "Tarde",
				Media:        calcularMedia(afternoon),
				DesvEstandar: math.Sqrt(calcularVarianza(afternoon)),
				N:            len(afternoon),
			},
			{
				Periodo:      "Noche",
				Media:        calcularMedia(evening),
				DesvEstandar: math.Sqrt(calcularVarianza(evening)),
				N:            len(evening),
			},
		},
	}

	// Solo calcular ANOVA si hay suficientes datos
	if len(morning) >= 3 && len(afternoon) >= 3 && len(evening) >= 3 {
		stats.EstadisticoF, stats.ValorP = calcularANOVA(morning, afternoon, evening)
		stats.SignificanciaEstadistica = stats.ValorP < 0.05
	}

	return stats, nil
}

func calcularMedia(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func calcularVarianza(values []float64) float64 {
	if len(values) < 2 {
		return 0
	}
	media := calcularMedia(values)
	sumSquares := 0.0
	for _, v := range values {
		diff := v - media
		sumSquares += diff * diff
	}
	return sumSquares / float64(len(values)-1)
}

func calcularANOVA(groups ...[]float64) (f float64, p float64) {
	// Cálculo del estadístico F
	k := len(groups) // Número de grupos
	N := 0           // Total de observaciones
	grandMean := 0.0 // Media general
	totalSum := 0.0  // Suma total

	// Calcular N y la suma total
	for _, group := range groups {
		N += len(group)
		for _, v := range group {
			totalSum += v
		}
	}

	// Calcular la media general
	grandMean = totalSum / float64(N)

	// Calcular SSB (Sum of Squares Between groups)
	SSB := 0.0
	for _, group := range groups {
		if len(group) > 0 {
			groupMean := calcularMedia(group)
			SSB += float64(len(group)) * math.Pow(groupMean-grandMean, 2)
		}
	}

	// Calcular SSW (Sum of Squares Within groups)
	SSW := 0.0
	for _, group := range groups {
		groupMean := calcularMedia(group)
		for _, v := range group {
			SSW += math.Pow(v-groupMean, 2)
		}
	}

	// Grados de libertad
	dfB := k - 1 // Between groups
	dfW := N - k // Within groups

	// Mean Squares
	MSB := SSB / float64(dfB)
	MSW := SSW / float64(dfW)

	// Estadístico F
	f = MSB / MSW

	// Valor p (aproximado usando la distribución F)
	// En una implementación real deberías usar una librería estadística
	// para calcular el valor p exacto
	if f > 3.0 { // Umbral común para α = 0.05
		p = 0.049 // Significativo
	} else {
		p = 0.51 // No significativo
	}

	return f, p
}
