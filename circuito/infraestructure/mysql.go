package infraestructure

import (
	"database/sql"
	"second/circuito/domain"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db:db}
}

func (r *MySQLRepository) Save(circuito domain.Circuito) error {
	query := "INSERT INTO circuito(idMesa, color) VALUES (?, ?)"
	_,err := r.db.Exec(query, circuito.IdMesa, circuito.Color)
	if err != nil {
		return err
	}
	return err
}

func (r *MySQLRepository) GetAll() ([]domain.Circuito, error) {
	query := "SELECT idCircuito, idMesa, color FROM circuito"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var circuitos []domain.Circuito
	for rows.Next() {
		var circuito domain.Circuito
		if err := rows.Scan(&circuito.IdCircuito, &circuito.IdMesa, &circuito.Color); err != nil {
			return nil, err
		}
		circuitos = append(circuitos, circuito)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return circuitos, nil
}

func (r *MySQLRepository) GetById(idCircuito int) ([]domain.Circuito, error) {
	query := "SELECT idCircuito, idMesa, color FROM circuito WHERE idCircuito=?"
	rows, err := r.db.Query(query, idCircuito)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var robotStatus []domain.Circuito
	for rows.Next() {
		var circuito domain.Circuito
		if err := rows.Scan(&circuito.IdCircuito, &circuito.IdMesa, &circuito.Color); err != nil {
			return nil, err
		}
		robotStatus = append(robotStatus, circuito)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return robotStatus, nil
}

func (r *MySQLRepository) Delete(idCircuito int) error {
	query := "DELETE FROM circuito WHERE idCircuito=?"
	_, err := r.db.Exec(query, idCircuito)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLRepository) Update(idCircuito int, circuito domain.Circuito) error {
	query := "UPDATE circuito SET idMesa=?, color=? WHERE idCircuito=?"
	_,err := r.db.Exec(query, circuito.IdMesa, circuito.Color, idCircuito)


	if err != nil {
		return err
	}
	return err
}