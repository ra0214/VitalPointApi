package domain

type IBodyTemperature interface {
	SaveBodyTemperature(esp32ID string, tiempo string, temp_ambiente float64, temp_objeto float64) error
	GetAll() ([]BodyTemperature, error)
}

type BodyTemperature struct {
	ID           int32   `json:"id"`
	ESP32ID      string  `json:"esp32ID"`
	Timestamp    string  `json:"tiempo"`
	TempAmbiente float64 `json:"temp_ambiente"`
	TempObjeto   float64 `json:"temp_objeto"`
}

func NewBodyTemperature(esp32ID string, tiempo string, temp_ambiente float64, temp_objeto float64) *BodyTemperature {
	return &BodyTemperature{
		ESP32ID:      esp32ID,
		Timestamp:    tiempo,
		TempAmbiente: temp_ambiente,
		TempObjeto:   temp_objeto,
	}
}

func (bt *BodyTemperature) SetTemperature(temp_ambiente float64, temp_objeto float64) {
	bt.TempAmbiente = temp_ambiente
	bt.TempObjeto = temp_objeto
}
