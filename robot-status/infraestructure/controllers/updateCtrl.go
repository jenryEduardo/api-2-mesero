package controllers

import (
	"second/robot-status/application"
	"second/robot-status/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateRSCtrl struct {
	uc *application.UpdateRS
}

func NewUpdateRSCtrl (uc *application.UpdateRS) *UpdateRSCtrl {
	return &UpdateRSCtrl{uc:uc}
}

func (ctrl *UpdateRSCtrl) Run(c *gin.Context) {
	var rs domain.RobotStatus

	if err := c.ShouldBindJSON(&rs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error":"status es requerido"})
		return
	}

	id := c.Param("idEstado")
	idEstado, err := strconv.Atoi(id)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"El id debe ser númerico",
		})
		return
	}

	if err := ctrl.uc.Run(idEstado, rs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data": gin.H{
				"type": "Estado del robot",
				"attributes": gin.H{
					"idEstado": rs.IdEstado,
					"idRobot": rs.IdRobot,
					"Estado del robot": rs.Status,
				},
			},
		})
	}
	return
}