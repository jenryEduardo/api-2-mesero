package infraestructure

import (
	"database/sql"
	"second/robot-status/domain"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) Save(robotStatus domain.RobotStatus) error {
	query := "INSERT INTO estado_robot(idRobot, status) VALUES(?, ?)" 
	_, err := r.db.Exec(query, robotStatus.IdRobot, robotStatus.Status)
	if err != nil {
		return err
	}
	return err
}

func (r *MySQLRepository) GetAll() ([]domain.RobotStatus, error) {
	query := "SELECT  idRobot, status FROM estado_robot"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var robotStatus []domain.RobotStatus
	for rows.Next() {
		var rs domain.RobotStatus
		if err := rows.Scan(&rs.IdRobot, &rs.Status); err != nil {
			return nil, err
		}
		robotStatus = append(robotStatus, rs)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return robotStatus, nil
}

func (r *MySQLRepository) GetById(idEstado int) ([]domain.RobotStatus, error) {
	query := "SELECT idRobot, status FROM estado_robot WHERE idEstado=?"
	rows, err := r.db.Query(query, idEstado)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var robotStatus []domain.RobotStatus
	for rows.Next() {
		var rs domain.RobotStatus
		if err := rows.Scan( &rs.IdRobot, &rs.Status); err != nil {
			return nil, err
		}
		robotStatus = append(robotStatus, rs)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return robotStatus, nil
}

func (r *MySQLRepository) Delete(idEstado int) error {
	query := "DELETE FROM estado_robot WHERE idEstado=?"
	_, err := r.db.Exec(query, idEstado)
	if err != nil {
		return err
	}
	return nil
} 

func (r *MySQLRepository) Update(idEstado int, RS domain.RobotStatus) error {
	query := "UPDATE estado_robot SET status=? WHERE idEstado=?"
	_,err := r.db.Exec(query, RS.Status, idEstado)
	if err != nil {
		return err
	}
	return err
}