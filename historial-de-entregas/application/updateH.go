package application

import "second/historial-de-entregas/domain"

type UpdateHistorial struct {
	repo domain.IHistorial
}

func NewUpdateHistorial(repo domain.IHistorial) *UpdateHistorial {
	return &UpdateHistorial{repo:repo}
}

func (c *UpdateHistorial) Run(idHistorial int, historial domain.Historial) error {
	return c.repo.Update(idHistorial, historial)
}