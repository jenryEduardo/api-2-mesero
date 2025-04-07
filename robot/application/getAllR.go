package application

import "second/robot/domain"

type GetAllRobot struct {
	repo domain.IRobot
}

func NewGetAllRobot(repo domain.IRobot) *GetAllRobot {
	return &GetAllRobot{repo: repo}
}

func (c *GetAllRobot) Run() ([]domain.Robot, error) {
	return c.repo.GetAll()
}
