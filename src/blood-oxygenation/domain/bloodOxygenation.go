package domain

type IBloodOxygenation interface {
	SaveBloodOxygenation(esp32ID string, bloodOxygenation float64) error
	GetAll() ([]BloodOxygenation, error)
}

type BloodOxygenation struct {
	ID int32  `json:"id"`
	ESP32ID     string  `json:"esp32ID"`
	BloodOxygenation float64 `json:"bloodOxygenation"`
}

func NewBloodOxygenation(esp32ID string, bloodOxygenation float64) *BloodOxygenation {
	return &BloodOxygenation{
		ESP32ID:         esp32ID,
		BloodOxygenation: bloodOxygenation,
	}
}

func (bt *BloodOxygenation) SetBloodOxygenation(bloodOxygenation float64) {
	bt.BloodOxygenation = bloodOxygenation
}