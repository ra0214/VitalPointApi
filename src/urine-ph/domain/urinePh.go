package domain

type IUrinePh interface {
	SaveUrinePh(esp32ID string, timestamp string, ph int32) error
	GetAll() ([]UrinePh, error)
}

type UrinePh struct {
	ID int32  `json:"id"`
	ESP32ID     string  `json:"esp32ID"`
	Timestamp string  `json:"tiempo"`
	PH int32   `json:"ph"`
}

func NewUrinePh(esp32ID string, tiempo string, ph int32) *UrinePh {
	return &UrinePh{
		ESP32ID:  esp32ID,
		Timestamp: tiempo,
		PH:       ph,
	}
}

func (ur *UrinePh) SetUrinePh(ph int32) {
	ur.PH = ph
}