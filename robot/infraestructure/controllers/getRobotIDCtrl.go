package controllers

import (
	"second/robot/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetRobotIDCtrl struct {
	uc *application.GetByIdRobot
}

func NewGetRobotByIDCtrl(uc *application.GetByIdRobot) *GetRobotIDCtrl {
	return &GetRobotIDCtrl{uc: uc}
}

func (ctrl *GetRobotIDCtrl) Run(c *gin.Context) {
	id := c.Param("idRobot")
	idRobot, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	robots, err := ctrl.uc.Run(idRobot)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": robots,
		})
	}
}