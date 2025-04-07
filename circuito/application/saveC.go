package application

import "second/circuito/domain"

type SaveCircuito struct {
	repo domain.ICircuito
}

func NewSaveCircuito(repo domain.ICircuito) *SaveCircuito {
	return &SaveCircuito{repo:repo}
}

func (c *SaveCircuito) Run(circuito domain.Circuito) error {
	return c.repo.Save(circuito)
}