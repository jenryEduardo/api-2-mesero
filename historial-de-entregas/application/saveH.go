package application

import "second/historial-de-entregas/domain"

type SaveHistorial struct {
	repo domain.IHistorial
}

func NewSaveHistorial(repo domain.IHistorial) *SaveHistorial {
	return &SaveHistorial{repo:repo}
}

func(c *SaveHistorial) Run(historial domain.Historial) error {
	return c.repo.Save(historial)
}