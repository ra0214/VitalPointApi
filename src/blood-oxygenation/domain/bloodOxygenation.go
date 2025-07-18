package domain

import "database/sql"

type IBloodOxygenation interface {
	SaveBloodOxygenation(esp32ID string, timestamp string, ir int32, red int32, spo2 float64) error
	GetAll() ([]BloodOxygenation, error)
}

type BloodOxygenation struct {
	ID        int32           `json:"id"`
	ESP32ID   string          `json:"esp32ID"`
	Timestamp string          `json:"tiempo"`
	IR        int32           `json:"ir"`
	Red       int32           `json:"red"`
	SpO2      sql.NullFloat64 `json:"spo2"`
}

// Si quieres seguir usando tus constructores:
func NewBloodOxygenation(esp32ID string, tiempo string, ir int32, red int32, spo2 float64) *BloodOxygenation {
	return &BloodOxygenation{
		ESP32ID:   esp32ID,
		Timestamp: tiempo,
		IR:        ir,
		Red:       red,
		SpO2:      sql.NullFloat64{Float64: spo2, Valid: true},
	}
}

func (bt *BloodOxygenation) SetBloodOxygenation(ir int32, red int32, spo2 float64) {
	bt.IR = ir
	bt.Red = red
	bt.SpO2 = sql.NullFloat64{Float64: spo2, Valid: true}
}
