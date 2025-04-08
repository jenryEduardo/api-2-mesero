package application

import "second/historial-de-entregas/domain"

type FindIdCircuito struct {
	repo domain.IHistorial
}

func NewFindIdCircuito(repo domain.IHistorial) *FindIdCircuito {
	return &FindIdCircuito{repo:repo}
}

func (c *FindIdCircuito) Run(idPedido int) (string,error) {
	return c.repo.FindIdCircuito(idPedido)
}