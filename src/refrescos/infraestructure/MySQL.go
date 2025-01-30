package infraestructure

import (
	"actividad/src/config"
	"actividad/src/refrescos/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IRefrescos = (*MySQL)(nil)

func NewMySQL() domain.IRefrescos {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}
func (mysql *MySQL) SaveRefrescos(marca string, precio float32) {
	query := "INSERT INTO refresco (marca, modelo, precio) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, marca, precio)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Refresco guardado correctamente: Marca: %s - Precio: %.2f", marca, precio)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
}

func (mysql *MySQL) GetAll() {
	query := "SELECT * FROM refresco"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	for rows.Next() {
		var id int
		var marca string
		var precio float32
		if err := rows.Scan(&id, &marca, &precio); err != nil {
			fmt.Printf("Error al escanear la fila: %v\n", err)
		}
		fmt.Printf("ID: %d, Marca: %s, Precio: %.2f\n", id, marca, precio)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterando sobre las filas: %v\n", err)
	}
}

func (mysql *MySQL) UpdateRefrescos(id int32, marca string, precio float32) error {
	query := "UPDATE refresco SET marca = ?, precio = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, marca, precio, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de actualización: %v", err)
	}
	return nil
}

func (mysql *MySQL) DeleteRefrescos(id int32) error {
	query := "DELETE FROM refresco WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de eliminación: %v", err)
	}
	return nil
}