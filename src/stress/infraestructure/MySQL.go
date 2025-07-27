package infraestructure

import (
	"fmt"
	"log"
	"vitalPoint/src/config"
	"vitalPoint/src/stress/domain"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IStress = (*MySQL)(nil)

// Constructor de la conexión a MySQL
func NewMySQL() domain.IStress {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

// Guardar el nivel de estrés en la base de datos
func (mysql *MySQL) SaveStress(esp32ID string, tiempo string, stress string) error {
	query := "INSERT INTO stress (esp32ID, tiempo, stress) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, esp32ID, tiempo, stress)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener las filas afectadas: %v", err)
	}

	if rowsAffected == 1 {
		log.Printf("[MySQL] - estrés guardado correctamente: Esp32ID:%s Estrés:%s", esp32ID, stress)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.Stress, error) {
	query := "SELECT id, esp32ID, tiempo, stress FROM stress"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var stresss []domain.Stress

	for rows.Next() {
		var stress domain.Stress
		if err := rows.Scan(&stress.ID, &stress.ESP32ID, &stress.Timestamp, &stress.Stress); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %v", err)
		}
		stresss = append(stresss, stress)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterando sobre las filas: %v", err)
	}
	return stresss, nil
}

// Obtener la última temperatura para un ESP32
func (mysql *MySQL) GetLatestTemperature(esp32ID string) (float64, error) {
	query := "SELECT stress FROM bodyTemperature WHERE esp32ID = ? ORDER BY tiempo DESC LIMIT 1"
	rows, err := mysql.conn.FetchRows(query, esp32ID)
	if err != nil {
		return 0, fmt.Errorf("Error al obtener la última temperatura: %v", err)
	}
	defer rows.Close()

	var temperature float64
	if rows.Next() {
		if err := rows.Scan(&temperature); err != nil {
			return 0, fmt.Errorf("Error al escanear temperatura: %v", err)
		}
	}
	return temperature, nil
}

// Obtener la última oxigenación para un ESP32
func (mysql *MySQL) GetLatestOxygenation(esp32ID string) (float64, error) {
	query := "SELECT spo2 FROM bloodOxygenation WHERE esp32ID = ? ORDER BY tiempo DESC LIMIT 1"
	rows, err := mysql.conn.FetchRows(query, esp32ID)
	if err != nil {
		return 0, fmt.Errorf("Error al obtener la última oxigenación: %v", err)
	}
	defer rows.Close()

	var oxygenation float64
	if rows.Next() {
		if err := rows.Scan(&oxygenation); err != nil {
			return 0, fmt.Errorf("Error al escanear oxigenación: %v", err)
		}
	}
	return oxygenation, nil
}
