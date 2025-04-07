package dependencies

import (
	"second/core"
	"second/historial-de-entregas/application"
	"second/historial-de-entregas/infraestructure"
	"second/historial-de-entregas/infraestructure/controllers"
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

func SaveHistorial() *controllers.SaveHistorialCtrl {
	ucSaveH := application.NewSaveHistorial(&mySQL)
	return controllers.NewSaveHistorialCtrl(ucSaveH)
}

func GetAllHistorial() *controllers.GetAllHistorialCtrl {
	ucGetAllH := application.NewGetAllHistorial(&mySQL)
	return controllers.NewGetAllHistorialCtrl(ucGetAllH)
}

func GetHByID() *controllers.GetHistorialByIDCtrl {
	ucGetById := application.NewGetHistorialByID(&mySQL)
	return controllers.NewGetHistorialByIDCtrl(ucGetById)
}

func UpdateHistorial() *controllers.UpdateHistorialCtrl {
	ucUpdateH := application.NewUpdateHistorial(&mySQL)
	return controllers.NewUpdateHistorialCtrl(ucUpdateH)
}

func DeleteHistorial() *controllers.DeleteHistorialCtrl {
	ucDeleteH := application.NewDeleteHistorial(&mySQL)
	return controllers.NewDeleteHistorialCtrl(ucDeleteH)
}

func FindCircuito() *controllers.FindIdCircuitoCtrl {
	ucFindCircuito := application.NewFindIdCircuito(&mySQL)
	return controllers.NewFindIdCircuitoCtrl(ucFindCircuito)
}