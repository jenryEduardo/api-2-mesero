package application

import "second/robot/domain"

type UpdateRobot struct {
	repo domain.IRobot
}

func NewUpdateRobot(repo domain.IRobot) *UpdateRobot {
	return &UpdateRobot{repo: repo}
}

func (c *UpdateRobot) Run(idRobot int, robot domain.Robot) error {
	return c.repo.Update(idRobot, robot)
}