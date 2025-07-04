package domain

type IBloodOxygenation interface {
	SaveBloodOxygenation(esp32ID string, bloodOxygenation float64, timestamp string, ir int32, red int32) error
	GetAll() ([]BloodOxygenation, error)
}

type BloodOxygenation struct {
	ID int32  `json:"id"`
	ESP32ID     string  `json:"esp32ID"`
	BloodOxygenation float64 `json:"bloodOxygenation"`
	Timestamp string  `json:"tiempo"`
	IR int32   `json:"ir"`
	Red int32  `json:"red"`
}

func NewBloodOxygenation(esp32ID string, bloodOxygenation float64, tiempo string, ir int32, red int32) *BloodOxygenation {
	return &BloodOxygenation{
		ESP32ID:         esp32ID,
		BloodOxygenation: bloodOxygenation,
		Timestamp:       tiempo,
		IR:              ir,
		Red:             red,
	}
}

func (bt *BloodOxygenation) SetBloodOxygenation(bloodOxygenation float64) {
	bt.BloodOxygenation = bloodOxygenation
}