package domain

// Primero definimos la interfaz completa
type IUrinePh interface {
	SaveUrinePh(esp32ID string, tiempo string, ph float64) error
	GetAll() ([]UrinePh, error)
	GetStats() (*UrinePhStats, error)
}

type UrinePh struct {
	ID        int32   `json:"id"`
	ESP32ID   string  `json:"esp32ID"`
	Timestamp string  `json:"tiempo"`
	PH        float64 `json:"ph"`
}

type UrinePhStats struct {
	// Estadísticas básicas
	Media        float64 `json:"media"`
	Mediana      float64 `json:"mediana"`
	DesvEstandar float64 `json:"desviacion_estandar"`
	Minimo       float64 `json:"minimo"`
	Maximo       float64 `json:"maximo"`

	// Datos para ANOVA
	GruposHorarios []struct {
		Periodo  string  `json:"periodo"`
		Media    float64 `json:"media"`
		Varianza float64 `json:"varianza"`
		N        int     `json:"n"`
	} `json:"grupos_horarios"`

	// Resultados ANOVA
	EstadisticoF             float64 `json:"estadistico_f"`
	ValorP                   float64 `json:"valor_p"`
	SignificanciaEstadistica bool    `json:"significancia_estadistica"`
}

func NewUrinePh(esp32ID string, tiempo string, ph float64) *UrinePh {
	return &UrinePh{
		ESP32ID:   esp32ID,
		Timestamp: tiempo,
		PH:        ph,
	}
}

func (ur *UrinePh) SetUrinePh(ph float64) {
	ur.PH = ph
}
