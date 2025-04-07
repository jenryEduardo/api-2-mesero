package application

import "second/circuito/domain"

type GetAllCircuitos struct {
	repo domain.ICircuito
}

func NewGetAllCircuitos(repo domain.ICircuito) *GetAllCircuitos {
	return &GetAllCircuitos{repo:repo}
}

func(c *GetAllCircuitos) Run() ([]domain.Circuito, error) {
	return c.repo.GetAll()
}