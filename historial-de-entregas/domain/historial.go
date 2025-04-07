package domain

type Historial struct {
	IdHistorial int 		`json:"id_historial"`
	IdPedido int			`json:"idPedido"`
	IdCircuito int 
	IdRobot int 			`json:"idRobot"`
	Estatus_entrega string 	`json:"estatus_entrega"`
}