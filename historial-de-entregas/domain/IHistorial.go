package domain

type IHistorial interface {
	Save(historial Historial) error
	GetAll()([]Historial, error)
	GetById(idHistorial int)([]Historial,error)
	Update(idHistorial int, historial Historial) error
	Delete(idHistorial int) error 
	FindIdCircuito(idPedido int) (int,error)
}