package application

import "second/robot-status/domain"

type UpdateRS struct {
	repo domain.IRobotStatus
}

func NewUpdateRS(repo domain.IRobotStatus) *UpdateRS {
	return &UpdateRS{repo: repo}
}

func (c *UpdateRS) Run(idEstado int, RS domain.RobotStatus) error {
	return c.repo.Update(idEstado, RS)
}