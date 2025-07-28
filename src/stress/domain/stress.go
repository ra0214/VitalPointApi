package domain

type IStress interface {
	SaveStress(esp32ID string, timestamp string, stress string) error
	GetAll() ([]Stress, error)
	GetLatestTemperature(esp32ID string) (float64, error)
	GetLatestOxygenation(esp32ID string) (float64, error)
	GetLatestStress(esp32ID string) (*Stress, error)
	// Nuevo método para obtener datos de correlación
	GetCorrelationData(esp32ID string) ([]StressCorrelation, error)
}

// Nueva estructura para datos de correlación
type StressCorrelation struct {
	ESP32ID     string  `json:"esp32_id"`
	Temperatura float64 `json:"temperatura"`
	Oxigenacion float64 `json:"oxigenacion"`
	Stress      string  `json:"stress"`
	Timestamp   string  `json:"timestamp"`
}

type Stress struct {
	ID        int32  `json:"id"`
	ESP32ID   string `json:"esp32ID"`
	Timestamp string `json:"tiempo"`
	Stress    string `json:"stress"`
}

func NewStress(esp32ID string, tiempo string, stress string) *Stress {
	return &Stress{
		ESP32ID:   esp32ID,
		Timestamp: tiempo,
		Stress:    stress,
	}
}

func (ur *Stress) SetStress(stress string) {
	ur.Stress = stress
}
