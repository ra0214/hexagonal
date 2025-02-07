package infraestructure

import (
	"actividad/src/config"
	"actividad/src/tiendas/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.ITienda = (*MySQL)(nil)

func NewMySQL() domain.ITienda {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) SaveTienda(nombre string, ubicacion string) {
	query := "INSERT INTO tienda (nombre, ubicacion) VALUES ( ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, nombre, ubicacion)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Tienda guardada correctamente: Nombre: %s - Ubicacion: %s", nombre, ubicacion)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
}

func (mysql *MySQL) GetAll() ([]domain.Tienda, error) {
    query := "SELECT * FROM tienda"
    rows := mysql.conn.FetchRows(query)
    defer rows.Close()

    var tiendas []domain.Tienda

    for rows.Next() {
        var tienda domain.Tienda
        if err := rows.Scan(&tienda.ID, &tienda.Nombre, &tienda.Ubicacion); err != nil {
            return nil, fmt.Errorf("error scanning row: %w", err)
        }
        tiendas = append(tiendas, tienda)
    }

    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error iterating rows: %w", err)
    }

    return tiendas, nil
}

func (mysql *MySQL) UpdateTienda(id int32, nombre string, ubicacion string) error {
	query := "UPDATE tienda SET nombre = ?, ubicacion = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, nombre, ubicacion, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de actualización: %v", err)
	}
	return nil
}

func (mysql *MySQL) DeleteTienda(id int32) error {
	query := "DELETE FROM tienda WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de eliminación: %v", err)
	}
	return nil
}