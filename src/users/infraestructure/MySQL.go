package infraestructure

import (
	"fmt"
	"log"
	"vitalPoint/src/config"
	"vitalPoint/src/users/domain"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IUser = (*MySQL)(nil)

func NewMySQL() domain.IUser {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) SaveUser(userName string, email string, password string, role string) error {
	query := "INSERT INTO user (userName, email, password, role) VALUES (?, ?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, userName, email, password, role)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Usuario creado correctamente: Username:%s Email:%s Role:%s", userName, email, role)
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.User, error) {
	query := "SELECT id, userName, email, password, role FROM user"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.UserName, &user.Email, &user.Password, &user.Role); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterando sobre las filas: %v", err)
	}
	return users, nil
}

func (mysql *MySQL) UpdateUser(id int32, userName string, email string, password string, role string) error {
	query := "UPDATE user SET userName = ?, email = ?, password = ?, role = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, userName, email, password, role, id)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Usuario actualizado correctamente: ID:%d Username:%s Email:%s Role:%s", id, userName, email, role)
	}
	return nil
}

func (mysql *MySQL) DeleteUser(id int32) error {
	query := "DELETE FROM user WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Usuario eliminado correctamente: ID: %d", id)
	} else {
		log.Println("[MySQL] - No se eliminó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetUserByCredentials(userName string) (*domain.User, error) {
	query := "SELECT id, userName, email, password FROM user WHERE userName = ?"
	row := mysql.conn.FetchRow(query, userName) // No retorna error aquí

	var user domain.User
	err := row.Scan(&user.ID, &user.UserName, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("usuario no encontrado")
	}

	return &user, nil
}
