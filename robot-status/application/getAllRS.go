package application

import "second/robot-status/domain"

type GetAllRS struct {
	repo domain.IRobotStatus
}

func NewGetAllRS(repo domain.IRobotStatus) *GetAllRS {
	return &GetAllRS{repo: repo}
}

func (c *GetAllRS) Run() ([]domain.RobotStatus, error) {
	return c.repo.GetAll()
}