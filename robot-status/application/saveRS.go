package application

import "second/robot-status/domain"

type SaveRS struct {
	repo domain.IRobotStatus
}

func NewSaveRS(repo domain.IRobotStatus) *SaveRS {
	return &SaveRS{repo: repo}
}

func (c *SaveRS) Run(rs domain.RobotStatus) error {
	return c.repo.Save(rs)
}
