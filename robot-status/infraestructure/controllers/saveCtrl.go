package controllers

import (
	"second/robot-status/application"
	"second/robot-status/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaveRSCtrl struct {
	uc *application.SaveRS
}

func NewSaveRSCtrl(uc *application.SaveRS) *SaveRSCtrl {
	return &SaveRSCtrl{uc: uc}
}

func (ctrl *SaveRSCtrl) Run(c *gin.Context) {
	var rs domain.RobotStatus
	
	if err := c.ShouldBindJSON(&rs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return 
	}

	err := ctrl.uc.Run(rs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"status": true,
			"data": gin.H{
				"type": "robot-status",
				"attributes": gin.H{
					"idRobot": rs.IdRobot,
					"estado": rs.Status,
				},
			},
		})
	}
}