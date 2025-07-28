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
	Media              float64          `json:"media"`
	Mediana            float64          `json:"mediana"`
	Moda               float64          `json:"moda"`
	DesviacionEstandar float64          `json:"desviacion_estandar"`
	Minimo             float64          `json:"minimo"`
	Maximo             float64          `json:"maximo"`
	FrecuenciaData     []FrecuenciaData `json:"frecuenciaData"`
}

type FrecuenciaData struct {
	Valor      float64 `json:"valor"`
	Frecuencia int     `json:"frecuencia"`
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
