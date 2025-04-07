package infraestructure

import (
	"database/sql"
	"second/robot/domain"
)

type MySQLRepository struct {
	db *sql.DB
}

func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) Save(robot domain.Robot) error {
	query := "INSERT INTO robot(idRobot, alias) VALUES (?,?)"
	_, err := r.db.Exec(query,robot.IdRobot, robot.Alias )
	if err != nil {
		return err
	}
	return err
}

func (r *MySQLRepository) GetAll() ([]domain.Robot, error) {
	query:= "SELECT idRobot, alias FROM robot"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var robots []domain.Robot
	for rows.Next() {
		var robot domain.Robot
		if err := rows.Scan(&robot.IdRobot, &robot.Alias); err != nil {
			return nil, err
		}
		robots = append(robots, robot)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return robots, nil
}

func (r *MySQLRepository) Update(id int, robot domain.Robot) error {
	query := "UPDATE robot SET alias=? WHERE idRobot=?"
	_, err := r.db.Exec(query, robot.Alias, id)
	if err != nil {
		return err
	}
	return err
}

func (r *MySQLRepository) Delete(id int) error {
	query := "DELETE FROM robot WHERE idRobot=?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *MySQLRepository) GetById(id int) ([]domain.Robot, error) {
	query := "SELECT idRobot, alias FROM robot WHERE idRobot=?"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var robots []domain.Robot
	for rows.Next() {
		var robot domain.Robot
		if err := rows.Scan(&robot.IdRobot, &robot.Alias); err != nil {
			return nil, err
		}
		robots = append(robots, robot)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return robots, nil
}


