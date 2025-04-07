package application

import "second/robot/domain"

type DeleteRobot struct {	
	repo domain.IRobot
}

func NewDeleteRobot(repo domain.IRobot) *DeleteRobot {
	return &DeleteRobot{repo: repo}
}

func (c *DeleteRobot) Run(idRobot int) error {
	return c.repo.Delete(idRobot)
}