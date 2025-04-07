package routes

import (
	"github.com/gin-gonic/gin"
	"second/pedido/infraestructure/controllers"
)

func SetUpRoutes(routes *gin.Engine){
	router:=routes.Group("/pedidos")

	router.GET("/:idPedido",controllers.GetIdPedido)
}







