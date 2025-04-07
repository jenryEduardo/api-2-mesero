package application

import "second/historial-de-entregas/domain"

type DeleteHistorial struct {
	repo domain.IHistorial
}

func NewDeleteHistorial (repo domain.IHistorial) *DeleteHistorial {
	return &DeleteHistorial{repo:repo}
}

func (c *DeleteHistorial) Run(idHistorial int) error {
	return c.repo.Delete(idHistorial)
}