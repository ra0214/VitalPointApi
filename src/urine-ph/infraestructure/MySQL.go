package infraestructure

import (
	"fmt"
	"log"
	"vitalPoint/src/config"
	"vitalPoint/src/urine-ph/application" // Cambiamos este import
	"vitalPoint/src/urine-ph/domain"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IUrinePh = (*MySQL)(nil)

// Constructor de la conexión a MySQL
func NewMySQL() domain.IUrinePh {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

// Guardar el pH de la orina en la base de datos
func (mysql *MySQL) SaveUrinePh(esp32ID string, tiempo string, ph float64) error {
	query := "INSERT INTO urineph (esp32ID, tiempo, ph) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, esp32ID, tiempo, ph)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener las filas afectadas: %v", err)
	}

	if rowsAffected == 1 {
		log.Printf("[MySQL] - pH de la orina guardado correctamente: Esp32ID:%s pH:%f", esp32ID, float64(ph))
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.UrinePh, error) {
	query := "SELECT id, esp32ID, tiempo, ph FROM urineph"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var urinePhs []domain.UrinePh

	for rows.Next() {
		var urinePh domain.UrinePh
		if err := rows.Scan(&urinePh.ID, &urinePh.ESP32ID, &urinePh.Timestamp, &urinePh.PH); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %v", err)
		}
		log.Printf("Dato leído: ID=%d, ESP32ID=%s, Tiempo=%s, PH=%.2f",
			urinePh.ID, urinePh.ESP32ID, urinePh.Timestamp, urinePh.PH)
		urinePhs = append(urinePhs, urinePh)
	}

	log.Printf("Total de registros leídos: %d", len(urinePhs))

	return urinePhs, nil
}

// Implementar el nuevo método GetStats
func (mysql *MySQL) GetStats() (*domain.UrinePhStats, error) {
	// Obtener todos los registros
	readings, err := mysql.GetAll()
	if err != nil {
		return nil, fmt.Errorf("error al obtener los datos: %v", err)
	}

	// Crear el analizador y pasarle los datos directamente
	analyzer := application.NewAnalyzeUrinePh()

	// Ejecutar el análisis pasando los datos
	stats, err := analyzer.Execute(readings)
	if err != nil {
		return nil, fmt.Errorf("error al analizar los datos: %v", err)
	}

	return stats, nil
}
