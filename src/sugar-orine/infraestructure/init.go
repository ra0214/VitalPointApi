package infraestructure

import (
	"vitalPoint/src/config"
	"log"
)

func InitUrinePh() error {
	log.Println("Inicializando datos...")

	db, err := config.GetDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	log.Println("Conexión a la base de datos para glucosa establecida correctamente")
	return nil
}