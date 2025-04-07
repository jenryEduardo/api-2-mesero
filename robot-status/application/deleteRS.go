package application

import "second/robot-status/domain"

type DeleteRS struct {
	repo domain.IRobotStatus
}

func NewDeleteRS(repo domain.IRobotStatus) *DeleteRS {
	return &DeleteRS{repo: repo}
}

func (c *DeleteRS) Run(idEstado int) error {
	return c.repo.Delete(idEstado)
}