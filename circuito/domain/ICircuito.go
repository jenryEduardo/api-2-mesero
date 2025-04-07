package domain

type ICircuito interface {
	Save(circuito Circuito) error
	GetAll()([]Circuito,error)
	GetById(idCircuito int) ([]Circuito, error)
	Update(idCircuito int, circuito Circuito) error
	Delete(idCircuito int) error
}