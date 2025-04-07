package domain

type IRobotStatus interface {
	Save(RS RobotStatus) error
	GetAll() ([]RobotStatus, error)
	GetById(idEstado int) ([]RobotStatus, error)
	Delete(idEstado int) error
	Update(idEstado int, RS RobotStatus) error 
}