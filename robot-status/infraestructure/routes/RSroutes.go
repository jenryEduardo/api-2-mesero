package routes

import (
	"second/robot-status/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func RSroutes(router *gin.Engine) {
	routes := router.Group(("/rs"))

	saveRS := dependencies.SaveRS()
	getAllRS := dependencies.GetAllRS()
	getByIDRS := dependencies.GetRSByID()
	deleteRS := dependencies.DeleteRS()
	updateRS := dependencies.UpdateRS()

	routes.POST("/", saveRS.Run)
	routes.GET("/", getAllRS.Run)
	routes.GET(":idEstado", getByIDRS.Run)
	routes.DELETE("/:idEstado", deleteRS.Run)
	routes.PUT("/:idEstado", updateRS.Run)
}