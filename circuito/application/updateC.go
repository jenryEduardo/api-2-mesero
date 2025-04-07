package application

import "second/circuito/domain"

type UpdateCircuito struct {
	repo domain.ICircuito
}

func NewUpdateCircuito(repo domain.ICircuito) *UpdateCircuito {
	return &UpdateCircuito{repo:repo}
}

func (c *UpdateCircuito) Run(idCircuito int, circuito domain.Circuito) error {
	return c.repo.Update(idCircuito, circuito)
} 