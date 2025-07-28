package application

import (
	"fmt"
	"math"
	"sort"
	"vitalPoint/src/body-temperature/domain"
)

type TempStats struct {
	Media         float64   `json:"media"`
	Mediana       float64   `json:"mediana"`
	Moda          float64   `json:"moda"`
	DesvEstandar  float64   `json:"desviacion_estandar"`
	Minimo        float64   `json:"minimo"`
	Maximo        float64   `json:"maximo"`
	FrecRelativa  []float64 `json:"frecuencia_relativa"`
	FrecAcumulada []float64 `json:"frecuencia_acumulada"`
	Intervalos    []string  `json:"intervalos"`
}

type AnalyzeBodyTemperature struct {
	repo domain.IBodyTemperature
}

func NewAnalyzeBodyTemperature(repo domain.IBodyTemperature) *AnalyzeBodyTemperature {
	return &AnalyzeBodyTemperature{repo: repo}
}

func (a *AnalyzeBodyTemperature) Execute() (*TempStats, error) {
	// Obtener todos los datos
	temps, err := a.repo.GetAll()
	if err != nil {
		return nil, err
	}

	// Extraer solo las temperaturas objeto
	values := make([]float64, len(temps))
	for i, t := range temps {
		values[i] = t.TempObjeto
	}

	// Ordenar valores para cálculos estadísticos
	sort.Float64s(values)

	stats := &TempStats{
		Media:        calcularMedia(values),
		Mediana:      calcularMediana(values),
		Moda:         calcularModa(values),
		DesvEstandar: calcularDesvEstandar(values),
		Minimo:       values[0],
		Maximo:       values[len(values)-1],
	}

	// Calcular frecuencias
	intervalos, frecRelativa, frecAcumulada := calcularFrecuencias(values, 10) // 10 intervalos
	stats.Intervalos = intervalos
	stats.FrecRelativa = frecRelativa
	stats.FrecAcumulada = frecAcumulada

	return stats, nil
}

func calcularMedia(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}

func calcularMediana(values []float64) float64 {
	n := len(values)
	if n%2 == 0 {
		return (values[n/2-1] + values[n/2]) / 2
	}
	return values[n/2]
}

func calcularModa(values []float64) float64 {
	frecMap := make(map[float64]int)
	for _, v := range values {
		frecMap[v]++
	}

	maxFrec := 0
	moda := 0.0
	for val, frec := range frecMap {
		if frec > maxFrec {
			maxFrec = frec
			moda = val
		}
	}
	return moda
}

func calcularDesvEstandar(values []float64) float64 {
	media := calcularMedia(values)
	sum := 0.0
	for _, v := range values {
		sum += math.Pow(v-media, 2)
	}
	return math.Sqrt(sum / float64(len(values)))
}

func calcularFrecuencias(values []float64, numIntervalos int) ([]string, []float64, []float64) {
	min, max := values[0], values[len(values)-1]
	rango := max - min
	amplitud := rango / float64(numIntervalos)

	intervalos := make([]string, numIntervalos)
	frecRelativa := make([]float64, numIntervalos)
	frecAcumulada := make([]float64, numIntervalos)

	n := float64(len(values))

	// Calcular frecuencias
	for i := 0; i < numIntervalos; i++ {
		limInf := min + float64(i)*amplitud
		limSup := limInf + amplitud
		intervalos[i] = fmt.Sprintf("%.1f-%.1f", limInf, limSup)

		// Contar frecuencias
		count := 0
		for _, v := range values {
			if v >= limInf && v < limSup {
				count++
			}
		}

		frecRelativa[i] = float64(count) / n * 100
		if i == 0 {
			frecAcumulada[i] = frecRelativa[i]
		} else {
			frecAcumulada[i] = frecAcumulada[i-1] + frecRelativa[i]
		}
	}

	return intervalos, frecRelativa, frecAcumulada
}
