package domain

type IBloodOxygenation interface {
	SaveBloodOxygenation(esp32ID string, timestamp string, ir int32, red int32) error
	GetAll() ([]BloodOxygenation, error)
}

type BloodOxygenation struct {
	ID int32  `json:"id"`
	ESP32ID     string  `json:"esp32ID"`
	Timestamp string  `json:"tiempo"`
	IR int32   `json:"ir"`
	Red int32  `json:"red"`
}

func NewBloodOxygenation(esp32ID string, tiempo string, ir int32, red int32) *BloodOxygenation {
	return &BloodOxygenation{
		ESP32ID:  esp32ID,
		Timestamp: tiempo,
		IR:       ir,
		Red:             red,
	}
}

func (bt *BloodOxygenation) SetBloodOxygenation(ir int32, red int32) {
	bt.IR = ir
	bt.Red = red
}