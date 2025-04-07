package controllers

import (
	"second/robot/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteRobotCtrl struct {
	uc *application.DeleteRobot
}

func NewDeleteRobotCtrl(uc *application.DeleteRobot) *DeleteRobotCtrl {
	return &DeleteRobotCtrl{uc: uc}
}

func (ctrl *DeleteRobotCtrl) Run(c *gin.Context) {
	id := c.Param("idRobot")
	idRobot, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El id debe ser númerico"})
		return
	}

	err = ctrl.uc.Run(idRobot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": "Robot eliminado",
		})
	}
}