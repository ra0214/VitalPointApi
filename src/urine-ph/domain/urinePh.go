package domain

type IUrinePh interface {
	SaveUrinePh(esp32ID string, timestamp string, ph float64) error
	GetAll() ([]UrinePh, error)
}

type UrinePh struct {
	ID int32  `json:"id"`
	ESP32ID     string  `json:"esp32ID"`
	Timestamp string  `json:"tiempo"`
	PH float64   `json:"ph"`
}

func NewUrinePh(esp32ID string, tiempo string, ph float64) *UrinePh {
	return &UrinePh{
		ESP32ID:  esp32ID,
		Timestamp: tiempo,
		PH:       ph,
	}
}

func (ur *UrinePh) SetUrinePh(ph float64) {
	ur.PH = ph
}