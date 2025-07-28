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
	// Usar temp_objeto que parece ser la temperatura corporal
	query := "SELECT temp_objeto FROM bodytemp WHERE esp32_id = ? ORDER BY tiempo DESC LIMIT 1"
	fmt.Printf("Ejecutando query temperatura: %s con esp32_id: %s\n", query, esp32ID)

	rows, err := mysql.conn.FetchRows(query, esp32ID)
	if err != nil {
		fmt.Printf("Error en query temperatura: %v\n", err)
		return 0, fmt.Errorf("Error al obtener la última temperatura: %v", err)
	}
	defer rows.Close()

	var temperature float64
	if rows.Next() {
		if err := rows.Scan(&temperature); err != nil {
			fmt.Printf("Error escaneando temperatura: %v\n", err)
			return 0, fmt.Errorf("Error al escanear temperatura: %v", err)
		}
		fmt.Printf("Temperatura encontrada: %.2f\n", temperature)
	} else {
		fmt.Printf("No se encontraron datos de temperatura para esp32_id: %s\n", esp32ID)
	}
	return temperature, nil
}

// Obtener la última oxigenación para un ESP32
func (mysql *MySQL) GetLatestOxygenation(esp32ID string) (float64, error) {
	// Cambiar a bloodoxygenation y esp32ID
	query := "SELECT spo2 FROM bloodoxygenation WHERE esp32ID = ? AND spo2 IS NOT NULL ORDER BY tiempo DESC LIMIT 1"
	fmt.Printf("Ejecutando query oxigenación: %s con esp32ID: %s\n", query, esp32ID)

	rows, err := mysql.conn.FetchRows(query, esp32ID)
	if err != nil {
		fmt.Printf("Error en query oxigenación: %v\n", err)
		return 0, fmt.Errorf("Error al obtener la última oxigenación: %v", err)
	}
	defer rows.Close()

	var oxygenation float64
	if rows.Next() {
		if err := rows.Scan(&oxygenation); err != nil {
			fmt.Printf("Error escaneando oxigenación: %v\n", err)
			return 0, fmt.Errorf("Error al escanear oxigenación: %v", err)
		}
		fmt.Printf("Oxigenación encontrada: %.2f\n", oxygenation)
	} else {
		fmt.Printf("No se encontraron datos de oxigenación para esp32ID: %s\n", esp32ID)
	}
	return oxygenation, nil
}

func (mysql *MySQL) GetCorrelationData(esp32ID string) ([]domain.StressCorrelation, error) {
	query := `
        SELECT bt.temp_objeto, bo.spo2, s.stress, s.tiempo
        FROM stress s
        INNER JOIN bodytemp bt ON DATE(bt.tiempo) = DATE(s.tiempo) 
            AND HOUR(bt.tiempo) = HOUR(s.tiempo) 
            AND MINUTE(bt.tiempo) = MINUTE(s.tiempo)
        INNER JOIN bloodoxygenation bo ON DATE(bo.tiempo) = DATE(s.tiempo)
            AND HOUR(bo.tiempo) = HOUR(s.tiempo)
            AND MINUTE(bo.tiempo) = MINUTE(s.tiempo)
        WHERE s.esp32ID = ?
        ORDER BY s.tiempo DESC
        LIMIT 50`

	rows, err := mysql.conn.FetchRows(query, esp32ID)
	if err != nil {
		return nil, fmt.Errorf("Error obteniendo datos de correlación: %v", err)
	}
	defer rows.Close()

	var correlations []domain.StressCorrelation
	for rows.Next() {
		var c domain.StressCorrelation
		if err := rows.Scan(&c.Temperatura, &c.Oxigenacion, &c.Stress, &c.Timestamp); err != nil {
			return nil, fmt.Errorf("Error escaneando datos: %v", err)
		}
		correlations = append(correlations, c)
	}

	return correlations, nil
}
