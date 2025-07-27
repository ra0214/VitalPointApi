package application

// Calcula el nivel de estrés basado en temperatura y oxigenación
func CalcularNivelStress(temperatura float64, oxigenacion float64) string {
	if (temperatura > 38.5) || (oxigenacion < 90) {
		return "Alto"
	}
	if (temperatura >= 37.6 && temperatura <= 38.5) || (oxigenacion >= 90 && oxigenacion <= 94) {
		return "Medio"
	}
	if (temperatura >= 36 && temperatura <= 37.5) && (oxigenacion >= 95) {
		return "Bajo"
	}
	return "Medio" // Valor por defecto
}
