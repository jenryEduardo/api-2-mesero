package routes

import (
	"github.com/gin-gonic/gin"
	"second/ultrasonico/infra/controllers"
)

func SetUpRoutes(routes *gin.Engine){
	router:=routes.Group("/ultra")

	router.POST("/",controllers.SaveIn)
}