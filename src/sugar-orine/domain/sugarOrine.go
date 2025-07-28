package domain

type ISugarOrine interface {
	SaveSugarOrine(esp32ID string, timestamp string, glucosa string) error
	GetAll() ([]SugarOrine, error)
	GetStats() (*SugarOrineStats, error) // Nuevo método
}

type SugarOrine struct {
	ID        int32  `json:"id"`
	ESP32ID   string `json:"esp32ID"`
	Timestamp string `json:"tiempo"`
	Glucosa   string `json:"glucosa"`
}

type SugarOrineStats struct {
	FrecuenciaData []FrecuenciaData `json:"frecuenciaData"`
	Normal         float64          `json:"normal"`
	Moderado       float64          `json:"moderado"`
	Alto           float64          `json:"alto"`
}

type FrecuenciaData struct {
	Valor      string `json:"valor"`
	Frecuencia int    `json:"frecuencia"`
}

func NewSugarOrine(esp32ID string, tiempo string, glucosa string) *SugarOrine {
	return &SugarOrine{
		ESP32ID:   esp32ID,
		Timestamp: tiempo,
		Glucosa:   glucosa,
	}
}

func (ur *SugarOrine) SetSugarOrine(glucosa string) {
	ur.Glucosa = glucosa
}
