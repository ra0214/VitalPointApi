package domain

type IStress interface {
	SaveStress(esp32ID string, timestamp string, stress string) error
	GetAll() ([]Stress, error)
}

type Stress struct {
	ID int32  `json:"id"`
	ESP32ID     string  `json:"esp32ID"`
	Timestamp string  `json:"tiempo"`
	Stress string   `json:"stress"`
}

func NewStress(esp32ID string, tiempo string, stress string) *Stress {
	return &Stress{
		ESP32ID:  esp32ID,
		Timestamp: tiempo,
		Stress:       stress,
	}
}

func (ur *Stress) SetStress(stress string) {
	ur.Stress = stress
}