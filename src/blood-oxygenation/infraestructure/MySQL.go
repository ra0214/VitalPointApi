package infraestructure

import (
	"vitalPoint/src/config"
	"vitalPoint/src/blood-oxygenation/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IBloodOxygenation = (*MySQL)(nil)

// Constructor de la conexión a MySQL
func NewMySQL() domain.IBloodOxygenation {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

// Guardar la saturación de oxígeno en sangre en la base de datos
func (mysql *MySQL) SaveBloodOxygenation(esp32ID string, tiempo string, ir int32, red int32, spo2 float64) error {
	query := "INSERT INTO bloodoxygenation (esp32ID, tiempo, ir, red, spo2) VALUES (?, ?, ?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, esp32ID, tiempo, ir, red, spo2)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener las filas afectadas: %v", err)
	}

	if rowsAffected == 1 {
		log.Printf("[MySQL] - Saturación de oxígeno en sangre guardada correctamente: Esp32ID:%s Saturación:%f", esp32ID, float64(ir+red)/2.0, spo2)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.BloodOxygenation, error) {
	query := "SELECT id, esp32ID, tiempo, ir, red, spo2 FROM bloodoxygenation"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var bloodOxygens []domain.BloodOxygenation

	for rows.Next() {
		var bloodOxygen domain.BloodOxygenation
		if err := rows.Scan(&bloodOxygen.ID, &bloodOxygen.ESP32ID, &bloodOxygen.Timestamp, &bloodOxygen.IR, &bloodOxygen.Red, &bloodOxygen.SpO2); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %v", err)
		}
		bloodOxygens = append(bloodOxygens, bloodOxygen)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterando sobre las filas: %v", err)
	}
	return bloodOxygens, nil
}
