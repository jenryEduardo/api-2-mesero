package application

import "second/robot-status/domain"

type GetbyIdRS struct {
	repo domain.IRobotStatus
}

func NewGetbyIdRS(repo domain.IRobotStatus) *GetbyIdRS {
	return &GetbyIdRS{repo: repo}
}

func (c *GetbyIdRS) Run(idEstado int) ([]domain.RobotStatus, error) {
	return c.repo.GetById(idEstado)
}