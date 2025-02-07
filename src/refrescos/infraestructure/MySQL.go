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
	query := "INSERT INTO refresco (marca, precio) VALUES (?, ?)"
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

func (mysql *MySQL) GetAll() ([]domain.Refrescos, error){
	query := "SELECT * FROM refresco"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var refrescos []domain.Refrescos

	for rows.Next() {
		var refresco domain.Refrescos
		if err := rows.Scan(&refresco.ID, &refresco.Marca, &refresco.Precio); err != nil {
            return nil, fmt.Errorf("error scanning row: %w", err)
		}
		refrescos = append(refrescos, refresco)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return refrescos, nil
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