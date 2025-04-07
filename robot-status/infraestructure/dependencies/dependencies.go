package dependencies

import (
	"second/core"
	"second/robot-status/application"
	"second/robot-status/infraestructure"
	"second/robot-status/infraestructure/controllers"
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

func SaveRS() *controllers.SaveRSCtrl {
	ucSaveRS := application.NewSaveRS(&mySQL)
	return controllers.NewSaveRSCtrl(ucSaveRS)
}

func GetAllRS() *controllers.GetAllRSCtrl {
	ucGetAllRS := application.NewGetAllRS(&mySQL)
	return controllers.NewGetAllRSCtrl(ucGetAllRS)
}

func GetRSByID() *controllers.GetRsIDCtrl {
	ucGetRSByID := application.NewGetbyIdRS(&mySQL)
	return controllers.NewGetRsIDCtrl(ucGetRSByID)
}

func DeleteRS() *controllers.DeleteRSCtrl {
	ucDeleteRS := application.NewDeleteRS(&mySQL)
	return controllers.NewDeleteRSCtrl(ucDeleteRS)
}

func UpdateRS() *controllers.UpdateRSCtrl {
	ucUpdateRS := application.NewUpdateRS(&mySQL)
	return controllers.NewUpdateRSCtrl(ucUpdateRS)
}