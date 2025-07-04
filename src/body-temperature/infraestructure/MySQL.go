package infraestructure

import (
	"vitalPoint/src/config"
	"vitalPoint/src/body-temperature/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IBodyTemperature = (*MySQL)(nil)

// Constructor de la conexión a MySQL
func NewMySQL() domain.IBodyTemperature {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

// Guardar la temperatura en la base de datos
func (mysql *MySQL) SaveBodyTemperature(esp32ID string, temperatura float64) error {
	query := "INSERT INTO bodytemp (esp32_id, temperatura) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, esp32ID, temperatura)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener las filas afectadas: %v", err)
	}

	if rowsAffected == 1 {
		log.Printf("[MySQL] - Temperatura guardada correctamente: Esp32ID:%s Temperatura:%f", esp32ID, temperatura)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.BodyTemperature, error) {
	query := "SELECT id, esp32_id, temperatura FROM bodytemp"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var bodyTemps []domain.BodyTemperature

	for rows.Next() {
		var bodyTemp domain.BodyTemperature
		if err := rows.Scan(&bodyTemp.ID, &bodyTemp.ESP32ID, &bodyTemp.Temperature); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %v", err)
		}
		bodyTemps = append(bodyTemps, bodyTemp)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterando sobre las filas: %v", err)
	}
	return bodyTemps, nil
}
