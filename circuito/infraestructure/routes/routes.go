package routes

import (
	"second/circuito/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func CircuitRoutes(router *gin.Engine) {
	routes := router.Group("/circuito")

	saveC := dependencies.SaveC()
	getAll := dependencies.GetAllC()
	getByID := dependencies.GetCByID()
	deleteC := dependencies.DeleteC()
	updateC := dependencies.UpdateC()

	routes.POST("/", saveC.Run)
	routes.GET("/", getAll.Run)
	routes.GET("/:idCircuito", getByID.Run)
	routes.DELETE("/:idCircuito", deleteC.Run)
	routes.PUT("/:idCircuito", updateC.Run)
}