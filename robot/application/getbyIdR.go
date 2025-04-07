package application

import "second/robot/domain"

type GetByIdRobot struct {
	repo domain.IRobot
}

func NewGetByIdRobot (repo domain.IRobot) *GetByIdRobot {
	return &GetByIdRobot{repo: repo}
}

func (c *GetByIdRobot) Run(id int) ([]domain.Robot, error) {
	return c.repo.GetById(id)
}