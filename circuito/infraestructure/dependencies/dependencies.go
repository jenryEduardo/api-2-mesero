package dependencies

import (
	"second/circuito/application"
	"second/circuito/infraestructure"
	"second/circuito/infraestructure/controllers"
	"second/core"
)

var (
	mySQL infraestructure.MySQLRepository
)

func Init() {
	db, err := core.InitMySQL()
	if err != nil {
		return
	}
	mySQL = *infraestructure.NewMySQLRepository(db)
}

func SaveC() *controllers.SaveCircuitoCtrl {
	ucSaveC := application.NewSaveCircuito(&mySQL)
	return controllers.NewSaveCircuitoCtrl(ucSaveC)
}

func GetAllC() *controllers.GetAllCircuitoCtrl {
	ucGetAllC := application.NewGetAllCircuitos(&mySQL)
	return controllers.NewGetAllCircuitosCtrl(ucGetAllC)
}

func GetCByID() *controllers.GetCircuitoByIdCtrl {
	ucGetByID := application.NewGetByIDCircuito(&mySQL)
	return controllers.NewGetByIDCircuitoCtrl(ucGetByID)
}

func UpdateC() *controllers.UpdateCircuitoCtrl {
	ucUpdateC := application.NewUpdateCircuito(&mySQL)
	return controllers.NewUpdateCircuitoCtrl(ucUpdateC)
}

func DeleteC() *controllers.DeleteCircuitoCtrl {
	ucDeleteC := application.NewDeleteCircuito(&mySQL)
	return controllers.NewDeleteCircuitoCtrl(ucDeleteC)
}
