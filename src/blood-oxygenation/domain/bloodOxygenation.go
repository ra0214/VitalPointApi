package domain

import (
	"database/sql"
	"encoding/json"
)

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
	SpO2      sql.NullFloat64 `json:"-"`
}

func (b BloodOxygenation) MarshalJSON() ([]byte, error) {
	type Alias BloodOxygenation
	return json.Marshal(&struct {
		SpO2 interface{} `json:"spo2"`
		Alias
	}{
		SpO2: func() interface{} {
			if b.SpO2.Valid {
				return b.SpO2.Float64
			}
			return nil
		}(),
		Alias: (Alias)(b),
	})
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
