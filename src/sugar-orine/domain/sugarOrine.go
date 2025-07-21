package domain

type ISugarOrine interface {
	SaveSugarOrine(esp32ID string, timestamp string, glucosa string) error
	GetAll() ([]SugarOrine, error)
}

type SugarOrine struct {
	ID int32  `json:"id"`
	ESP32ID     string  `json:"esp32ID"`
	Timestamp string  `json:"tiempo"`
	Glucosa string  `json:"glucosa"`
}

func NewSugarOrine(esp32ID string, tiempo string, glucosa string) *SugarOrine {
	return &SugarOrine{
		ESP32ID:  esp32ID,
		Timestamp: tiempo,
		Glucosa:  glucosa,
	}
}

func (ur *SugarOrine) SetSugarOrine(glucosa string) {
	ur.Glucosa = glucosa
}