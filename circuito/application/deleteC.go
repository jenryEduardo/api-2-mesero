package application

import "second/circuito/domain"

type DeleteCircuito struct {
	repo domain.ICircuito	
}

func NewDeleteCircuito(repo domain.ICircuito) *DeleteCircuito {
	return &DeleteCircuito{repo:repo}
}

func (c *DeleteCircuito) Run(idCircuito int) error {
	return c.repo.Delete(idCircuito)
}