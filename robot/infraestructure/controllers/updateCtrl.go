package controllers

import (
	"second/robot/application"
	"second/robot/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateRobotCtrl struct {
	uc *application.UpdateRobot
}

func NewUpdateRobotCtrl(uc *application.UpdateRobot) *UpdateRobotCtrl {
	return &UpdateRobotCtrl{uc: uc}
}

func (ctrl *UpdateRobotCtrl) Run(c *gin.Context)  {
	var robot domain.Robot

	if err := c.ShouldBindJSON(&robot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Todos los campos son requeridos"})
		return 
	}

	id := c.Param("idRobot")
	idRobot, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "El id del robot debe ser un número"})
			return 
	}

	if err := ctrl.uc.Run(idRobot, robot); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": gin.H{
				"type": "robot",	
				"idRobot": robot.IdRobot,
				"attributes": gin.H{
					"alias": robot.Alias,
				},
			},
		})
	} 
	return 
}