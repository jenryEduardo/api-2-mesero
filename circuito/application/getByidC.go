package application

import "second/circuito/domain"

type GetCircuitoById struct {
	repo domain.ICircuito
}

func NewGetByIDCircuito(repo domain.ICircuito) *GetCircuitoById {
	return &GetCircuitoById{repo:repo}
}

func(c *GetCircuitoById) Run(idCircuito int) ([]domain.Circuito, error) {
	return c.repo.GetById(idCircuito)
}