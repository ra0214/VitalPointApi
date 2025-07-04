package domain

type IBodyTemperature interface {
	SaveBodyTemperature(esp32ID string, temperature float64) error
	GetAll() ([]BodyTemperature, error)
}

type BodyTemperature struct {
	ID int32  `json:"id"`
	ESP32ID     string  `json:"esp32ID"`
	Temperature float64 `json:"temperature"`
}

func NewBodyTemperature(esp32ID string, temperature float64) *BodyTemperature {
	return &BodyTemperature{
		ESP32ID:     esp32ID,
		Temperature: temperature,
	}
}

func (bt *BodyTemperature) SetTemperature(temperature float64) {
	bt.Temperature = temperature
}