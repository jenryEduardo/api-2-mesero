package domain

type IRobot interface {
	Save(robot Robot) error
	Update(idRobot int, robot Robot) error
	Delete(idRobot int) error
	GetById(idRobot int) ([]Robot, error)
	GetAll() ([]Robot, error)
}