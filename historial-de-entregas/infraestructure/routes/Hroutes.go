package routes

import (
	"second/historial-de-entregas/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func HistorialRoutes(router *gin.Engine) {
	routes := router.Group("/historial")

	SaveH := dependencies.SaveHistorial()
	GetAllH := dependencies.GetAllHistorial()
	GetById := dependencies.GetHByID()
	UpdateH := dependencies.UpdateHistorial()
	DeleteH := dependencies.DeleteHistorial()
	FindCircuito := dependencies.FindCircuito()

	routes.POST("/", SaveH.Run)
	routes.GET("/", GetAllH.Run)
	routes.GET("/:id_historial", GetById.Run)
	routes.PUT("/:id_historial", UpdateH.Run)
	routes.DELETE("/id_historial", DeleteH.Run)
	routes.GET("/circuito/:idPedido", FindCircuito.Run)
}