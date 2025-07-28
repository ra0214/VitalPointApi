package application

import (
	"fmt"
	"math"
)

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
	frecuencias := make(map[float64]int)
	for _, v := range values {
		frecuencias[v]++
	}

	maxFrec := 0
	var moda float64
	for valor, frec := range frecuencias {
		if frec > maxFrec {
			maxFrec = frec
			moda = valor
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
	amplitud := (max - min) / float64(numIntervalos)

	intervalos := make([]string, numIntervalos)
	frecRelativa := make([]float64, numIntervalos)
	frecAcumulada := make([]float64, numIntervalos)

	total := float64(len(values))

	for i := 0; i < numIntervalos; i++ {
		limInf := min + float64(i)*amplitud
		limSup := limInf + amplitud
		intervalos[i] = fmt.Sprintf("%.1f-%.1f", limInf, limSup)

		// Contar frecuencias para niveles de SpO2
		for _, v := range values {
			if v >= limInf && (v < limSup || (i == numIntervalos-1 && v <= limSup)) {
				frecRelativa[i]++
			}
		}

		// Convertir a porcentaje
		frecRelativa[i] = (frecRelativa[i] / total) * 100

		// Calcular frecuencia acumulada
		if i == 0 {
			frecAcumulada[i] = frecRelativa[i]
		} else {
			frecAcumulada[i] = frecAcumulada[i-1] + frecRelativa[i]
		}
	}

	return intervalos, frecRelativa, frecAcumulada
}

func calcularPorcentajesRango(values []float64) []float64 {
	total := float64(len(values))
	rangos := make([]float64, 4)

	for _, v := range values {
		switch {
		case v >= 95:
			rangos[0]++ // Normal (95-100%)
		case v >= 90:
			rangos[1]++ // Leve (90-94%)
		case v >= 85:
			rangos[2]++ // Moderado (85-89%)
		default:
			rangos[3]++ // Severo (<85%)
		}
	}

	// Convertir a porcentajes
	for i := range rangos {
		rangos[i] = (rangos[i] / total) * 100
	}

	return rangos
}
