package routes

import (
	"second/robot/infraestructure/dependencies"

	"github.com/gin-gonic/gin"
)

func RobotRoutes(router *gin.Engine) {
	routes := router.Group("/robot")

	SaveRobot := dependencies.SaveRobot()
	GetRobotByID := dependencies.GetRobotByID()
	GetAllRobots := dependencies.GetAllRobots()
	UpdateRobot := dependencies.UpdateRobot()
	DeleteRobot := dependencies.DeleteRobot()

	routes.POST("/", SaveRobot.Run)
	routes.GET("/:idRobot", GetRobotByID.Run)
	routes.GET("/", GetAllRobots.Run)
	routes.PUT("/:idRobot", UpdateRobot.Run)
	routes.DELETE("/:idRobot", DeleteRobot.Run)
}