package infraestructure

import (
	"fmt"
	"log"
	"vitalPoint/src/config"
	"vitalPoint/src/sugar-orine/domain"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.ISugarOrine = (*MySQL)(nil)

// Constructor de la conexión a MySQL
func NewMySQL() domain.ISugarOrine {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

// Guardar el nivel de glucosa en la base de datos
func (mysql *MySQL) SaveSugarOrine(esp32ID string, tiempo string, glucosa string) error {
	query := "INSERT INTO sugarorine (esp32ID, tiempo, glucosa) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, esp32ID, tiempo, glucosa)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error al obtener las filas afectadas: %v", err)
	}

	if rowsAffected == 1 {
		log.Printf("[MySQL] - nivel de glucosa guardado correctamente: Esp32ID:%s Glucosa:%s", esp32ID, glucosa)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.SugarOrine, error) {
	query := "SELECT id, esp32ID, tiempo, glucosa FROM sugarorine"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var sugarOrines []domain.SugarOrine

	for rows.Next() {
		var sugarOrine domain.SugarOrine
		if err := rows.Scan(&sugarOrine.ID, &sugarOrine.ESP32ID, &sugarOrine.Timestamp, &sugarOrine.Glucosa); err != nil {
			return nil, fmt.Errorf("Error al escanear la fila: %v", err)
		}
		sugarOrines = append(sugarOrines, sugarOrine)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterando sobre las filas: %v", err)
	}
	return sugarOrines, nil
}
