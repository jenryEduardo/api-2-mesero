package infraestructure

import (
	"database/sql"
	"log"
	"second/historial-de-entregas/domain"
)


type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db:db}
}

func (r *MySQLRepository) Save(historial domain.Historial) error {
	query:= "INSERT INTO historial_entrega(idPedido,idCircuito, idRobot, estatus_entrega) VALUES(?,?,?,?)"
	_, err := r.db.Exec(query, historial.IdPedido,historial.IdCircuito, historial.IdRobot, historial.Estatus_entrega)
	if err != nil {
		return err
	}
	return err
}

func (r *MySQLRepository) GetAll()([]domain.Historial, error) {
	query:= "SELECT id_historial, idPedido,idCircuito, idRobot, estatus_entrega FROM historial_entrega"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	} 
	defer rows.Close()

	var historial []domain.Historial

	for rows.Next() {
		var h domain.Historial
		if err := rows.Scan(&h.IdHistorial, &h.IdPedido,&h.IdCircuito, &h.IdRobot, &h.Estatus_entrega); err != nil {
			return nil, err
		}
		historial = append(historial, h)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return historial, nil
}

func (r *MySQLRepository) GetById(idHistorial int) ([]domain.Historial, error) {
	query := "SELECT id_historial, idPedido, idCircuito, idRobot, estatus_entrega FROM historial_entrega WHERE id_historial=?"

	rows, err := r.db.Query(query, idHistorial)
	if err != nil {
		return nil,err
	}

	defer rows.Close()

	var historial []domain.Historial

	for rows.Next() {
		var h domain.Historial

		if err := rows.Scan(&h.IdHistorial, &h.IdPedido,&h.IdCircuito, &h.IdRobot, &h.Estatus_entrega); err != nil {
			return nil, err
		}
		historial = append(historial, h)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return historial, nil
}

func (r *MySQLRepository) Update(idHistorial int, historial domain.Historial) error {
	query := "UPDATE historial_entrega SET idPedido=?, idCircuito=?, idRobot=?, estatus_entrega=? WHERE id_historial=?"
	_, err := r.db.Exec(query, historial.IdPedido,historial.IdCircuito, historial.IdRobot, historial.Estatus_entrega, idHistorial)
	if err != nil {
		return err
	}
	return err
}

func (r *MySQLRepository) Delete(idHistorial int) error{
	query := "DELETE FROM historial_entrega WHERE id_historial=?"
	_, err := r.db.Exec(query, idHistorial)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLRepository) FindIdCircuito(idPedido int) (string, error) {
	var idMesa int
	var idCircuito string

	// Primera consulta: obtener idMesa a partir de idPedido
	err := r.db.QueryRow("SELECT idMesa FROM pedido WHERE idPedido=?", idPedido).Scan(&idMesa)
	if err != nil {
		log.Println("Error al obtener idMesa:", err)
		return "", err
	}

	// Segunda consulta: obtener idCircuito a partir de idMesa
	err = r.db.QueryRow("SELECT color FROM circuito WHERE idMesa=?", idMesa).Scan(&idCircuito)
	if err != nil {
		log.Println("Error al obtener idCircuito:", err)
		return "", err
	}

	return idCircuito, nil
}