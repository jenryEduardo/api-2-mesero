package application

import "second/historial-de-entregas/domain"


type GetHistorialByID struct {
	repo domain.IHistorial
}

func NewGetHistorialByID(repo domain.IHistorial) *GetHistorialByID{
	return &GetHistorialByID{repo:repo}
}

func (c *GetHistorialByID) Run(idHistorial int) ([]domain.Historial, error) {
	return c.repo.GetById(idHistorial)
}